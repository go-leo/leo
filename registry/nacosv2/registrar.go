package nacosv2

import (
	"context"
	"errors"

	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"

	"codeup.aliyun.com/qimao/leo/leo/internal/stringx"
	"codeup.aliyun.com/qimao/leo/leo/registry"
)

type Registrar struct {
	namingClient naming_client.INamingClient
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
	return nil
}

func (r *Registrar) Deregister(ctx context.Context, service registry.ServiceInstance) error {
	param := r.toDeregisterInstance(ctx, service)
	_, err := r.namingClient.DeregisterInstance(param)
	if err != nil {
		return err
	}
	return nil
}

func (r *Registrar) toRegisterInstance(ctx context.Context, service registry.ServiceInstance) vo.RegisterInstanceParam {
	params := registry.ParamsFromContext(ctx)
	weight, ok := params["Weight"].(float64)
	if !ok {
		weight = 1.0
	}
	enable, ok := params["Enable"].(bool)
	if !ok {
		enable = true
	}
	healthy, ok := params["Healthy"].(bool)
	if !ok {
		healthy = true
	}
	clusterName, _ := params["ClusterName"].(string)
	groupName, _ := params["GroupName"].(string)
	ephemeral, ok := params["Ephemeral"].(bool)
	if !ok {
		ephemeral = true
	}
	param := vo.RegisterInstanceParam{
		Ip:          service.IP(),           // 服务实例IP
		Port:        uint64(service.Port()), // 服务实例port
		Weight:      weight,                 // 权重
		Enable:      enable,                 // 是否上线
		Healthy:     healthy,                // 是否健康
		Metadata:    service.Metadata(),     // 扩展信息
		ClusterName: clusterName,
		ServiceName: service.Name(),
		GroupName:   groupName,
		Ephemeral:   ephemeral, // 是否临时实例
	}
	return param
}

func (r *Registrar) toDeregisterInstance(ctx context.Context, service registry.ServiceInstance) vo.DeregisterInstanceParam {
	params := registry.ParamsFromContext(ctx)
	clusterName, _ := params["ClusterName"].(string)
	groupName, _ := params["GroupName"].(string)
	ephemeral, ok := params["Ephemeral"].(bool)
	if !ok {
		ephemeral = true
	}
	param := vo.DeregisterInstanceParam{
		Ip:          service.IP(),           // 服务实例IP
		Port:        uint64(service.Port()), // 服务实例port
		Cluster:     clusterName,
		ServiceName: service.Name(),
		GroupName:   groupName,
		Ephemeral:   ephemeral, // 是否临时实例
	}
	return param
}

func NewRegistrar(factory NamingClientFactory) *Registrar {
	return &Registrar{namingClient: factory.Create()}
}
