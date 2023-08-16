package nacosv2

import (
	"context"
	"math"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

var _ registry.Discovery = new(Discovery)

type Discovery struct {
	namingClient naming_client.INamingClient
	nacosOptions *nacosOptions
}

func (d *Discovery) GetInstances(
	ctx context.Context,
	instance registry.ServiceInstance,
) ([]registry.ServiceInstance, error) {
	return d.getInstances(ctx, instance)
}

func (d *Discovery) GetServices(ctx context.Context) ([]string, error) {
	return d.getServices(ctx)
}

func (d *Discovery) Watch(ctx context.Context, instance registry.ServiceInstance) (registry.Watcher, error) {
	w := &Watcher{
		instance:  instance,
		discovery: d,
		closeC:    make(chan struct{}),
	}
	return w, w.init(ctx)
}

func (d *Discovery) getInstances(ctx context.Context, instance registry.ServiceInstance) ([]registry.ServiceInstance, error) {
	nacosServices, err := d.namingClient.SelectInstances(vo.SelectInstancesParam{
		Clusters:    d.nacosOptions.Clusters,
		ServiceName: instance.Name(),
		GroupName:   groupName(d.nacosOptions, instance),
		HealthyOnly: true,
	})
	if err != nil {
		return nil, err
	}
	serviceInstances := d.nacosHostsToServiceInstances(nacosServices, instance)
	return serviceInstances, nil
}

func (d *Discovery) nacosHostsToServiceInstances(hosts []model.Instance, serviceInstance registry.ServiceInstance) []registry.ServiceInstance {
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
		metadata := instance.Metadata
		var id string
		if len(metadata) > 0 {
			id = metadata[kInstanceId]
			delete(metadata, kInstanceId)
		}
		if stringx.IsBlank(id) {
			id = instance.InstanceId
		}
		theServiceInstance := registry.Builder().
			ID(id).
			Name(instance.ServiceName).
			IP(instance.Ip).
			Port(int(instance.Port)).
			Metadata(metadata).
			Scheme(serviceInstance.Scheme()).
			Build()
		serviceInstances = append(serviceInstances, theServiceInstance)
	}

	return serviceInstances
}

func (d *Discovery) getServices(ctx context.Context) ([]string, error) {
	servicesInfo, err := d.namingClient.GetAllServicesInfo(vo.GetAllServiceInfoParam{
		NameSpace: d.nacosOptions.Namespace,
		PageNo:    1,
		PageSize:  math.MaxInt32,
	})
	if err != nil {
		return nil, err
	}
	return servicesInfo.Doms, nil
}

func NewDiscovery(factory NamingClientFactoryFunc, opts ...NacosOption) (*Discovery, error) {
	namingClient, err := factory.Create()
	if err != nil {
		return nil, err
	}
	discovery := &Discovery{
		namingClient: namingClient,
		nacosOptions: &nacosOptions{
			Clusters:    []string{},
			ClusterName: "",
			Weight:      1.0,
			Namespace:   "",
			GroupName:   "",
		},
	}
	for _, opt := range opts {
		opt(discovery.nacosOptions)
	}
	return discovery, nil
}
