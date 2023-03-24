package nacosv2

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
)

type NamingClientFactory interface {
	Create() naming_client.INamingClient
}

type NamingClientFactoryFunc func() naming_client.INamingClient

func (f NamingClientFactoryFunc) Create() naming_client.INamingClient {
	return f()
}

type nacosOptions struct {
	clusters    []string
	clusterName string
	weight      float64
	nameSpace   string
}

type NacosOption func(r *nacosOptions)

func ClusterName(clusterName string) NacosOption {
	return func(r *nacosOptions) {
		r.clusterName = clusterName
	}
}

func Clusters(clusters []string) NacosOption {
	return func(r *nacosOptions) {
		r.clusters = clusters
	}
}

func Weight(weight float64) NacosOption {
	return func(r *nacosOptions) {
		r.weight = weight
	}
}

func NameSpace(nameSpace string) NacosOption {
	return func(r *nacosOptions) {
		r.nameSpace = nameSpace
	}
}
