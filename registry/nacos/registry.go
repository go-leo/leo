package nacos

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

	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"github.com/go-leo/stringx"

	"github.com/hmldd/leo/common/backoffx"
	"github.com/hmldd/leo/common/nacosx"
	"github.com/hmldd/leo/log"
	"github.com/hmldd/leo/registry"
)

var _ registry.Registrar = new(Registrar)

var _ registry.Discovery = new(Discovery)

const (
	_ID          = "id"
	_Transport   = "transport"
	_Version     = "version"
	_ClusterName = "cluster_name"
	_GroupName   = "group_name"
)

const Scheme = "nacos"

type Registrar struct {
	cli naming_client.INamingClient
}

func NewRegistrar(cli naming_client.INamingClient) *Registrar {
	return &Registrar{cli: cli}
}

func (r *Registrar) Register(ctx context.Context, service *registry.ServiceInfo) error {
	return r.register(service)
}

func (r *Registrar) Deregister(ctx context.Context, service *registry.ServiceInfo) error {
	return r.deregister(service)
}

func (r *Registrar) register(service *registry.ServiceInfo) error {
	metadata := make(map[string]string)
	for key, val := range service.Metadata {
		metadata[key] = val
	}
	metadata[_ID] = service.ID
	metadata[_Version] = service.Version
	metadata[_Transport] = service.Transport
	clusterName := r.extractString(metadata, _ClusterName)
	groupName := r.extractString(metadata, _GroupName)
	param := vo.RegisterInstanceParam{
		Ip:          service.Host,         //服务实例IP
		Port:        uint64(service.Port), //服务实例port
		Weight:      1.0,                  //权重
		Enable:      true,                 //是否上线
		Healthy:     true,                 //是否健康
		Metadata:    service.Metadata,     //扩展信息
		ClusterName: clusterName,          //集群名
		ServiceName: service.Name,
		GroupName:   groupName,
		Ephemeral:   false, //是否临时实例
	}
	_, err := r.cli.RegisterInstance(param)
	if err != nil {
		return err
	}
	return nil
}

func (r *Registrar) deregister(service *registry.ServiceInfo) error {
	metadata := make(map[string]string)
	for key, val := range service.Metadata {
		metadata[key] = val
	}
	clusterName := r.extractString(metadata, _ClusterName) //集群名称
	ephemeral := false                                     //是否临时实例
	port := uint64(service.Port)                           //服务实例port
	ip := service.Host                                     //服务实例IP
	serviceName := service.Name                            //服务名
	param := vo.DeregisterInstanceParam{
		Ip:          ip,
		Port:        port,
		Cluster:     clusterName,
		ServiceName: serviceName,
		GroupName:   _GroupName,
		Ephemeral:   ephemeral,
	}
	_, err := r.cli.DeregisterInstance(param)
	if err != nil {
		return err
	}
	return nil
}

func (r *Registrar) extractString(metadata map[string]string, key string) string {
	str := metadata[key]
	delete(metadata, key)
	return str
}

type Discovery struct {
	cli    naming_client.INamingClient
	logger log.Logger
	sync.RWMutex
	subscribeCallbackMap map[string]func(services []model.SubscribeService, err error)
	chanMap              map[string]chan []*registry.ServiceInfo
	cancelMap            map[string]context.CancelFunc
}

func NewDiscovery(cli naming_client.INamingClient, logger log.Logger) *Discovery {
	d := &Discovery{
		cli:                  cli,
		logger:               logger,
		RWMutex:              sync.RWMutex{},
		subscribeCallbackMap: make(map[string]func(services []model.SubscribeService, err error)),
		chanMap:              make(map[string]chan []*registry.ServiceInfo),
		cancelMap:            make(map[string]context.CancelFunc),
	}
	return d
}

func (d *Discovery) Scheme() string {
	return Scheme
}

func (d *Discovery) GetService(ctx context.Context, service *registry.ServiceInfo) ([]*registry.ServiceInfo, error) {
	return d.getService(service)
}

func (d *Discovery) Watch(ctx context.Context, service *registry.ServiceInfo) (<-chan []*registry.ServiceInfo, error) {
	return d.watch(service)
}

func (d *Discovery) StopWatch(ctx context.Context, service *registry.ServiceInfo) error {
	return d.stopWatch(service)
}

func (d *Discovery) getService(service *registry.ServiceInfo) ([]*registry.ServiceInfo, error) {
	if stringx.IsBlank(service.Name) {
		return nil, errors.New("service name is empty")
	}

	Clusters := d.getClusters(service)        //集群名称,多个集群用逗号分隔
	ServiceName := service.Name               //服务名
	GroupName := service.Metadata[_GroupName] //分组名
	param := vo.GetServiceParam{
		Clusters:    Clusters,
		ServiceName: ServiceName,
		GroupName:   GroupName,
	}
	nacosServices, err := d.cli.GetService(param)
	if err != nil {
		return nil, err
	}
	// for delete duplicate elements
	dde := make(map[string]*registry.ServiceInfo)
	for _, instance := range nacosServices.Hosts {
		if !instance.Enable {
			//忽略下线
			continue
		}
		if !instance.Healthy {
			//忽略不健康
			continue
		}
		if instance.Weight <= 0 {
			//忽略权重是负数
			continue
		}
		service := d.toServiceInfo(instance)
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

func (d *Discovery) watch(service *registry.ServiceInfo) (<-chan []*registry.ServiceInfo, error) {
	if stringx.IsBlank(service.Name) {
		return nil, fmt.Errorf("service %s's name is empty", service.Name)
	}
	d.Lock()
	defer d.Unlock()

	serviceKey := service.Name

	c, ok := d.chanMap[serviceKey]
	if ok {
		return nil, fmt.Errorf("service %s is watching", service.Name)
	}
	c = make(chan []*registry.ServiceInfo)
	ctx, cancelFunc := context.WithCancel(context.Background())
	d.cancelMap[serviceKey] = cancelFunc
	d.chanMap[serviceKey] = c
	d.subscribeCallbackMap[serviceKey] = func(services []model.SubscribeService, err error) {
		if err != nil {
			d.logger.Error(err)
		}
		serviceInfos, err := d.getService(service)
		if err != nil {
			d.logger.Error(err)
			return
		}
		c <- serviceInfos
	}

	param := vo.SubscribeParam{
		ServiceName:       service.Name,
		Clusters:          d.getClusters(service),
		GroupName:         service.Metadata[_GroupName],
		SubscribeCallback: d.subscribeCallbackMap[serviceKey],
	}
	if err := d.cli.Subscribe(&param); err != nil {
		close(c)
		return nil, err
	}
	go d.loop(ctx, service)
	return c, nil
}

func (d *Discovery) loop(ctx context.Context, service *registry.ServiceInfo) {
	delta := 10 * time.Millisecond
	serviceKey := service.Name
	for {
		sleepTime := backoffx.JitterUp(backoffx.Exponential(delta), 0.1)(context.Background(), 1)
		serviceInfos, err := d.getService(service)
		if err != nil {
			d.logger.Error(err)
		} else {
			d.chanMap[serviceKey] <- serviceInfos
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(sleepTime):
		}
	}
}

func (d *Discovery) stopWatch(service *registry.ServiceInfo) error {
	if stringx.IsBlank(service.Name) {
		return fmt.Errorf("service %s's name is empty", service.Name)
	}
	d.Lock()
	defer d.Unlock()
	serviceKey := service.Name

	// cancel looper
	cancelFunc, ok := d.cancelMap[serviceKey]
	if !ok {
		return fmt.Errorf("service %s isn't watching", service.Name)
	}
	cancelFunc()

	// Unsubscribe
	err := d.cli.Unsubscribe(&vo.SubscribeParam{
		ServiceName:       service.Name,
		Clusters:          d.getClusters(service),
		GroupName:         service.Metadata[_GroupName],
		SubscribeCallback: d.subscribeCallbackMap[serviceKey],
	})
	if err != nil {
		return err
	}

	// close channel
	close(d.chanMap[serviceKey])

	return nil
}

func (d *Discovery) getClusters(service *registry.ServiceInfo) []string {
	clusterName := service.Metadata[_ClusterName]
	var clusters []string
	if stringx.IsNotBlank(clusterName) {
		clusters = append(clusters, clusterName)
	}
	return clusters
}

func (d *Discovery) toServiceInfo(instance model.Instance) *registry.ServiceInfo {
	serviceInfo := &registry.ServiceInfo{
		ID:        instance.Metadata[_ID],
		Name:      instance.ServiceName,
		Transport: instance.Metadata[_Transport],
		Host:      instance.Ip,
		Port:      int(instance.Port),
		Metadata:  instance.Metadata,
		Version:   instance.Metadata[_Version],
	}
	return serviceInfo
}

type RegistrarFactory struct {
	URI *url.URL
}

func (factory *RegistrarFactory) Create() (registry.Registrar, error) {
	client, err := nacosx.NewNacosNamingClient(factory.URI)
	if err != nil {
		return nil, err
	}
	return NewRegistrar(client), nil
}

type DiscoveryFactory struct {
	URI    *url.URL
	Logger log.Logger
}

func (factory *DiscoveryFactory) Create() (registry.Discovery, error) {
	client, err := nacosx.NewNacosNamingClient(factory.URI)
	if err != nil {
		return nil, err
	}
	return NewDiscovery(client, factory.Logger), nil
}
