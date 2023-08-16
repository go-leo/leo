package nacosv2

import (
	"codeup.aliyun.com/qimao/leo/leo/registry"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

type NamingClientFactory interface {
	Create() (naming_client.INamingClient, error)
}

type NamingClientFactoryFunc func() (naming_client.INamingClient, error)

func (f NamingClientFactoryFunc) Create() (naming_client.INamingClient, error) {
	return f()
}

type nacosOptions struct {
	Clusters    []string
	ClusterName string
	Weight      float64
	Namespace   string
	GroupName   string
}

type NacosOption func(r *nacosOptions)

func ClusterName(clusterName string) NacosOption {
	return func(r *nacosOptions) {
		r.ClusterName = clusterName
	}
}

func Clusters(clusters []string) NacosOption {
	return func(r *nacosOptions) {
		r.Clusters = clusters
	}
}

func Weight(weight float64) NacosOption {
	return func(r *nacosOptions) {
		r.Weight = weight
	}
}

func NameSpace(nameSpace string) NacosOption {
	return func(r *nacosOptions) {
		r.Namespace = nameSpace
	}
}

func GroupName(name string) NacosOption {
	return func(r *nacosOptions) {
		r.GroupName = name
	}
}

func groupName(o *nacosOptions, instance registry.ServiceInstance) string {
	groupName := o.GroupName
	if len(groupName) <= 0 {
		groupName = instance.Scheme()
	}
	return groupName
}
