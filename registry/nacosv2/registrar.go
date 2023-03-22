package nacosv2

import (
	"context"
	"errors"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"github.com/go-leo/gox/stringx"

	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type Registrar struct {
	namingClient naming_client.INamingClient
	nacosOptions *nacosOptions
	deregisterC  chan struct{}
}

func (r *Registrar) Register(ctx context.Context, service registry.ServiceInstance) error {
	if stringx.IsBlank(service.ID()) {
		return errors.New("no service to register for nacos client")
	}
	param := r.toRegisterInstance(ctx, service)
	ok, err := r.namingClient.RegisterInstance(param)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("failed to register instance")
	}
	<-r.deregisterC
	return nil
}

func (r *Registrar) Deregister(ctx context.Context, service registry.ServiceInstance) error {
	defer close(r.deregisterC)
	param := r.toDeregisterInstance(ctx, service)
	ok, err := r.namingClient.DeregisterInstance(param)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("failed to deregister instance")
	}
	return nil
}

func (r *Registrar) toRegisterInstance(ctx context.Context, service registry.ServiceInstance) vo.RegisterInstanceParam {
	param := vo.RegisterInstanceParam{
		Ip:          service.IP(),           // 服务实例IP
		Port:        uint64(service.Port()), // 服务实例port
		Weight:      r.nacosOptions.weight,  // 权重
		Enable:      r.nacosOptions.enable,  // 是否上线
		Healthy:     r.nacosOptions.healthy, // 是否健康
		Metadata:    service.Metadata(),     // 扩展信息
		ClusterName: r.nacosOptions.clusterName,
		ServiceName: service.Name(),
		GroupName:   r.nacosOptions.groupName,
		Ephemeral:   r.nacosOptions.ephemeral, // 是否临时实例
	}
	return param
}

func (r *Registrar) toDeregisterInstance(ctx context.Context, service registry.ServiceInstance) vo.DeregisterInstanceParam {
	param := vo.DeregisterInstanceParam{
		Ip:          service.IP(),           // 服务实例IP
		Port:        uint64(service.Port()), // 服务实例port
		Cluster:     r.nacosOptions.clusterName,
		ServiceName: service.Name(),
		GroupName:   r.nacosOptions.groupName,
		Ephemeral:   r.nacosOptions.ephemeral, // 是否临时实例
	}
	return param
}

func NewRegistrar(factory NamingClientFactoryFunc, opts ...NacosOption) *Registrar {
	r := &Registrar{
		namingClient: factory.Create(),
		nacosOptions: &nacosOptions{
			clusterName: "",
			groupName:   "",
			weight:      1.0,
			healthy:     true,
			enable:      true,
			ephemeral:   true,
		},
		deregisterC: make(chan struct{}),
	}
	for _, opt := range opts {
		opt(r.nacosOptions)
	}
	return r
}
