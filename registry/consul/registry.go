package consul

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/sync/errgroup"

	"github.com/go-leo/stringx"

	"github.com/go-leo/backoffx"

	"github.com/go-leo/leo/v2/log"
	"github.com/go-leo/leo/v2/registry"
)

var _ registry.Registrar = new(Registrar)

var _ registry.Discovery = new(Discovery)

const Scheme = "consul"

const (
	_Transport = "transport"
	_Version   = "version"
	_Tag       = "tag"
)

type Registrar struct {
	cli             *api.Client
	healthCheckPath string
}

func NewRegistrar(cli *api.Client, healthCheckPath string) *Registrar {
	return &Registrar{cli: cli, healthCheckPath: healthCheckPath}
}

func (r *Registrar) Register(ctx context.Context, service *registry.ServiceInfo) error {
	return r.register(service.Clone())
}

func (r *Registrar) Deregister(ctx context.Context, service *registry.ServiceInfo) error {
	return r.deregister(service.Clone())
}

func (r *Registrar) register(service *registry.ServiceInfo) error {
	metadata := make(map[string]string)
	for key, val := range service.Metadata {
		metadata[key] = val
	}
	var tags []string
	if stringx.IsNotBlank(service.Version) {
		tags = append(tags, service.Version)
	}
	if stringx.IsNotBlank(service.Transport) {
		tags = append(tags, service.Transport)
	}
	if _tag, ok := metadata[_Tag]; ok {
		tags = append(tags, _tag)
		delete(metadata, _Tag)
	}
	interval := 5 * time.Second
	deregisterCriticalServiceAfter := time.Minute
	var GRPCChecks string
	var HTTPChecks string
	switch service.Transport {
	case registry.TransportGRPC:
		GRPCChecks = net.JoinHostPort(service.Host, strconv.Itoa(service.Port))
	case registry.TransportHTTP:
		HTTPChecks = fmt.Sprintf("%s://%s%s", "http", net.JoinHostPort(service.Host, strconv.Itoa(service.Port)), r.healthCheckPath)
	case registry.TransportHTTPS:
		HTTPChecks = fmt.Sprintf("%s://%s%s", "https", net.JoinHostPort(service.Host, strconv.Itoa(service.Port)), r.healthCheckPath)
	}
	asr := &api.AgentServiceRegistration{
		ID:      service.ID,
		Name:    service.Name,
		Meta:    metadata,
		Tags:    tags,
		Address: service.Host, //服务实例IP
		Port:    service.Port, //服务实例port
		Checks: api.AgentServiceChecks{
			{
				Interval:                       interval.String(),
				HTTP:                           HTTPChecks,
				GRPC:                           GRPCChecks,
				DeregisterCriticalServiceAfter: deregisterCriticalServiceAfter.String(),
			},
		},
	}
	err := r.cli.Agent().ServiceRegister(asr)
	if err != nil {
		return fmt.Errorf("failed register service, %w", err)
	}
	return nil
}

func (r *Registrar) deregister(service *registry.ServiceInfo) error {
	return r.cli.Agent().ServiceDeregister(service.ID)
}

type watchInfo struct {
	serviceInfoC chan []*registry.ServiceInfo
	watchPlan    *watch.Plan
	eg           *errgroup.Group
}

type Discovery struct {
	cli    *api.Client
	logger log.Logger
	sync.RWMutex
	watchInfoMap sync.Map
}

func NewDiscovery(cli *api.Client, logger log.Logger) *Discovery {
	d := &Discovery{
		cli:          cli,
		logger:       logger,
		RWMutex:      sync.RWMutex{},
		watchInfoMap: sync.Map{},
	}
	return d
}

func (d *Discovery) Scheme() string {
	return Scheme
}

func (d *Discovery) GetService(ctx context.Context, service *registry.ServiceInfo) ([]*registry.ServiceInfo, error) {
	return d.getService(ctx, service.Clone())
}

func (d *Discovery) Watch(ctx context.Context, service *registry.ServiceInfo) (<-chan []*registry.ServiceInfo, error) {
	return d.watch(ctx, service.Clone())
}

func (d *Discovery) StopWatch(ctx context.Context, service *registry.ServiceInfo) error {
	return d.stopWatch(ctx, service.Clone())
}

func (d *Discovery) getService(ctx context.Context, service *registry.ServiceInfo) ([]*registry.ServiceInfo, error) {
	if stringx.IsBlank(service.Name) {
		return nil, errors.New("service name is empty")
	}
	var tag string
	if stringx.IsNotBlank(service.Transport) {
		tag = service.Transport
	} else if stringx.IsNotBlank(service.Version) {
		tag = service.Version
	} else if _tag, ok := service.Metadata[_Tag]; ok {
		tag = _tag
	}
	serviceEntries, _, err := d.cli.Health().Service(service.Name, tag, true, &api.QueryOptions{})
	if err != nil {
		return nil, err
	}
	// for delete duplicate elements
	dde := make(map[string]*registry.ServiceInfo)
	for _, serviceEntry := range serviceEntries {
		service := d.toServiceInfo(serviceEntry)
		dde[net.JoinHostPort(service.Host, strconv.Itoa(service.Port))] = service
	}

	services := make([]*registry.ServiceInfo, 0, len(dde))
	for _, service := range dde {
		services = append(services, service)
	}

	// sort services by Address and Port
	sort.Slice(services, func(i, j int) bool {
		addressI := net.JoinHostPort(services[i].Host, strconv.Itoa(services[i].Port))
		addressJ := net.JoinHostPort(services[j].Host, strconv.Itoa(services[j].Port))
		return strings.Compare(addressI, addressJ) < 0
	})
	return services, nil
}

func (d *Discovery) watch(ctx context.Context, service *registry.ServiceInfo) (<-chan []*registry.ServiceInfo, error) {
	if stringx.IsBlank(service.Name) {
		return nil, fmt.Errorf("service %s's name is empty", service.Name)
	}
	d.Lock()
	defer d.Unlock()

	serviceKey := service.Name
	_, ok := d.watchInfoMap.Load(serviceKey)
	if ok {
		return nil, fmt.Errorf("service %s is watching", service.Name)
	}
	c := make(chan []*registry.ServiceInfo)

	watchPlan, err := watch.Parse(map[string]any{"type": "service", "service": service.Name})
	if err != nil {
		return nil, err
	}
	eg, ctx := errgroup.WithContext(ctx)
	info := watchInfo{
		eg:           eg,
		serviceInfoC: c,
		watchPlan:    watchPlan,
	}
	d.startWatch(ctx, info, service)
	d.loop(ctx, info, service)
	d.watchInfoMap.Store(serviceKey, info)
	return c, nil
}

func (d *Discovery) asyncGetServiceInfo(ctx context.Context, info watchInfo, service *registry.ServiceInfo) {
	info.eg.Go(func() error {
		serviceInfos, err := d.getService(ctx, service)
		if err != nil {
			d.logger.Error(err)
			return nil
		}
		info.serviceInfoC <- serviceInfos
		return nil
	})
}

func (d *Discovery) startWatch(ctx context.Context, info watchInfo, service *registry.ServiceInfo) {
	info.eg.Go(func() error {
		info.watchPlan.Handler = func(u uint64, i any) { d.asyncGetServiceInfo(ctx, info, service) }
		info.watchPlan.HybridHandler = func(u watch.BlockingParamVal, i any) { d.asyncGetServiceInfo(ctx, info, service) }
		return info.watchPlan.RunWithClientAndHclog(d.cli, hclog.Default())
	})
}

func (d *Discovery) loop(ctx context.Context, info watchInfo, service *registry.ServiceInfo) {
	info.eg.Go(func() error {
		delta := 10 * time.Second
		for {
			serviceInfos, err := d.getService(ctx, service)
			if err != nil {
				d.logger.Error(err)
			} else {
				info.serviceInfoC <- serviceInfos
			}
			sleepTime := backoffx.JitterUp(backoffx.Constant(delta), 0.1)(ctx, 0)
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(sleepTime):
			}
		}
	})
}

func (d *Discovery) stopWatch(_ context.Context, service *registry.ServiceInfo) error {
	if stringx.IsBlank(service.Name) {
		return fmt.Errorf("service %s's name is empty", service.Name)
	}
	d.Lock()
	defer d.Unlock()
	serviceKey := service.Name
	v, ok := d.watchInfoMap.Load(serviceKey)
	if !ok {
		return fmt.Errorf("service %s isn't watching", service.Name)
	}
	watchInfo, _ := v.(watchInfo)
	watchInfo.watchPlan.Stop()
	if err := watchInfo.eg.Wait(); err != nil {
		d.logger.Error(err)
	}
	close(watchInfo.serviceInfoC)
	return nil
}

func (d *Discovery) toServiceInfo(serviceEntry *api.ServiceEntry) *registry.ServiceInfo {
	metadata := make(map[string]string)
	for k, v := range serviceEntry.Service.Meta {
		metadata[k] = v
	}

	version := metadata[_Version]
	delete(metadata, _Version)

	transport := metadata[_Transport]
	delete(metadata, _Transport)

	serviceInfo := &registry.ServiceInfo{
		ID:        serviceEntry.Service.ID,
		Name:      serviceEntry.Service.Service,
		Transport: transport,
		Host:      serviceEntry.Service.Address,
		Port:      serviceEntry.Service.Port,
		Metadata:  metadata,
		Version:   version,
	}
	return serviceInfo
}

type RegistrarFactory struct {
	URI *url.URL
}

func (factory *RegistrarFactory) Create() (registry.Registrar, error) {
	client, err := NewClient(factory.URI)
	if err != nil {
		return nil, err
	}
	healthCheckPath := factory.URI.Query().Get("health_check_path")
	return NewRegistrar(client, healthCheckPath), nil
}

type DiscoveryFactory struct {
	URI    *url.URL
	Logger log.Logger
}

func (factory *DiscoveryFactory) Create() (registry.Discovery, error) {
	client, err := NewClient(factory.URI)
	if err != nil {
		return nil, err
	}
	return NewDiscovery(client, factory.Logger), nil
}

// consul://username:password@ip:port?scheme=http&datacenter=dev&token=12345&wait_time=1s&tls=true

func NewClient(uri *url.URL) (*api.Client, error) {
	query := uri.Query()
	config := &api.Config{
		Address:    uri.Host,
		Scheme:     query.Get("scheme"),
		Datacenter: query.Get("datacenter"),
		Token:      query.Get("token"),
		TokenFile:  query.Get("token_file"),
		TLSConfig: api.TLSConfig{
			Address:            "",
			CAFile:             "",
			CAPath:             "",
			CertFile:           "",
			KeyFile:            "",
			InsecureSkipVerify: false,
		},
	}

	if uri.User != nil && stringx.IsNotBlank(uri.User.Username()) {
		password, _ := uri.User.Password()
		config.HttpAuth = &api.HttpBasicAuth{Username: uri.User.Username(), Password: password}
	}

	waitTimeStr := query.Get("wait_time")
	if stringx.IsNotBlank(waitTimeStr) {
		duration, err := time.ParseDuration(waitTimeStr)
		if err != nil {
			return nil, fmt.Errorf("failed parse wait time %s, %w", waitTimeStr, err)
		}
		config.WaitTime = duration
	}

	if query.Get("tls") == "true" {
		config.TLSConfig = api.TLSConfig{
			Address:  query.Get("address"),
			CAFile:   query.Get("ca_file"),
			CAPath:   query.Get("ca_path"),
			CertFile: query.Get("cert_file"),
			KeyFile:  query.Get("key_file"),
		}
		if query.Get("insecure") == "true" {
			config.TLSConfig.InsecureSkipVerify = true
		}
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed new consul client, %w", err)
	}
	return client, nil
}
