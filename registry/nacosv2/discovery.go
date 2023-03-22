package nacosv2

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

var _ registry.Discovery = new(Discovery)

type Discovery struct {
	namingClient  naming_client.INamingClient
	instanceCache registry.Cache
	watchers      map[string][]chan []registry.ServiceInstance
	interval      time.Duration
	sync.RWMutex
}

func (d *Discovery) GetInstances(ctx context.Context, serviceName string) ([]registry.ServiceInstance, error) {
	serviceInstances, err := d.getInstancesFromNacos(ctx, serviceName)
	if err == nil && len(serviceInstances) > 0 {
		d.instanceCache.SetInstances(serviceName, serviceInstances)
		return serviceInstances, nil
	}
	instances, ok := d.instanceCache.GetInstances(serviceName)
	if ok && len(instances) > 0 {
		return instances, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, registry.ErrServiceNotFound
}

func (d *Discovery) GetServices(ctx context.Context) ([]string, error) {
	services, err := d.getServicesFromNacos(ctx)
	if err == nil && len(services) > 0 {
		d.instanceCache.SetServices(services)
		return services, nil
	}
	services = d.instanceCache.GetServices()
	if len(services) > 0 {
		return services, nil
	}
	if err != nil {
		return nil, err
	}
	return nil, registry.ErrServiceNotFound
}

func (d *Discovery) Watch(ctx context.Context, serviceName string) (<-chan []registry.ServiceInstance, error) {
	d.Lock()
	defer d.Unlock()
	ch := make(chan []registry.ServiceInstance)
	if _, ok := d.watchers[serviceName]; !ok {
		d.subscribe(ctx, serviceName)
		d.loop(ctx, serviceName)
	}
	d.watchers[serviceName] = append(d.watchers[serviceName], ch)
	return ch, nil
}

func (d *Discovery) Stop(ctx context.Context, serviceName string, watcher <-chan []registry.ServiceInstance) error {
	d.Lock()
	defer d.Unlock()
	// _, ok := d.instances[serviceName]
	// if !ok {
	// 	return ErrServiceNotFound
	// }
	// watchers, ok := d.watchers[serviceName]
	// if !ok {
	// 	return nil
	// }
	// var newWatchers []chan []ServiceInstance
	// for _, w := range watchers {
	// 	if w != watcher {
	// 		newWatchers = append(newWatchers, w)
	// 		continue
	// 	}
	// 	close(w)
	// }
	// d.watchers[serviceName] = newWatchers
	return nil
}

func (d *Discovery) onSetInstances(serviceName string, instances []registry.ServiceInstance) {
	d.Lock()
	defer d.Unlock()
	watchers, ok := d.watchers[serviceName]
	if !ok {
		return
	}
	for _, watcher := range watchers {
		select {
		case watcher <- instances:
		default:
			// 如果放不进去，跳过
			continue
		}
	}
}

func (d *Discovery) getInstancesFromNacos(ctx context.Context, serviceName string) ([]registry.ServiceInstance, error) {
	var clusters []string
	var groupName string
	params := registry.ParamsFromContext(ctx)
	clusters, _ = params["Clusters"].([]string)
	groupName, _ = params["GroupName"].(string)
	nacosServices, err := d.namingClient.GetService(vo.GetServiceParam{
		Clusters:    clusters,
		ServiceName: serviceName,
		GroupName:   groupName,
	})
	if err != nil {
		return nil, err
	}
	serviceInstances := d.nacosHostsToServiceInstances(nacosServices.Hosts)
	return serviceInstances, nil
}

func (d *Discovery) nacosHostsToServiceInstances(hosts []model.Instance) []registry.ServiceInstance {
	var serviceInstances []registry.ServiceInstance
	for _, instance := range hosts {
		if !instance.Enable {
			// 忽略下线
			continue
		}
		if !instance.Healthy {
			// 忽略不健康
			continue
		}
		if instance.Weight <= 0 {
			// 忽略权重是负数
			continue
		}
		serviceInstances = append(serviceInstances, d.nacosInstanceToServiceInstance(instance))
	}

	return serviceInstances
}

func (d *Discovery) nacosInstanceToServiceInstance(instance model.Instance) registry.ServiceInstance {
	return registry.NewServiceInstance(
		instance.InstanceId,
		instance.ServiceName,
		instance.Ip,
		int(instance.Port),
		instance.Metadata,
		instance.Metadata["scheme"],
	)
}

func (d *Discovery) getServicesFromNacos(ctx context.Context) ([]string, error) {
	var nameSpace string
	var groupName string
	params := registry.ParamsFromContext(ctx)
	nameSpace, _ = params["NameSpace"].(string)
	groupName, _ = params["GroupName"].(string)
	servicesInfo, err := d.namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: nameSpace,
		GroupName: groupName,
		PageNo:    1,
		PageSize:  math.MaxInt,
	})
	if err != nil {
		return nil, err
	}
	return servicesInfo.Doms, nil
}

func (d *Discovery) subscribe(ctx context.Context, serviceName string) error {
	var clusters []string
	var groupName string
	params := registry.ParamsFromContext(ctx)
	clusters, _ = params["Clusters"].([]string)
	groupName, _ = params["GroupName"].(string)
	err := d.namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName:       serviceName,
		Clusters:          clusters,
		GroupName:         groupName,
		SubscribeCallback: d.subscribeCallback(ctx, serviceName),
	})
	return err
}

func (d *Discovery) updateInstance(ctx context.Context, serviceName string) error {
	instances, err := d.GetInstances(ctx, serviceName)
	if err != nil {
		return err
	}
	d.RLock()
	watchers := d.watchers[serviceName]
	for _, watcher := range watchers {
		select {
		case watcher <- instances:
		default:
		}
	}
	d.RUnlock()
	return nil
}

func (d *Discovery) subscribeCallback(ctx context.Context, serviceName string) func(services []model.Instance, err error) {
	return func(services []model.Instance, err error) {
		if err != nil {
			fmt.Println("subscribeCallback:", err)
		}
		err = d.updateInstance(ctx, serviceName)
		if err != nil {
			fmt.Println("updateInstance:", err)
		}
	}
}

func (d *Discovery) loop(ctx context.Context, serviceName string) {
	interval := d.interval
	if interval <= 0 {
		interval = 10 * time.Second
	}
	for {
		err := d.updateInstance(ctx, serviceName)
		if err != nil {
			fmt.Println(err)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(interval):
		}
	}
}

func NewDiscovery(factory NamingClientFactory) *Discovery {
	discovery := &Discovery{
		namingClient:  factory.Create(),
		instanceCache: nil,
		watchers:      make(map[string][]chan []registry.ServiceInstance),
		interval:      0,
		RWMutex:       sync.RWMutex{},
	}
	// cache := NewServiceInstanceCache(discovery.onSetInstances)
	return discovery
}
