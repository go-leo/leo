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
	clusterName string
	groupName   string
	weight      float64
	healthy     bool
	enable      bool
	ephemeral   bool
}

type NacosOption func(r *nacosOptions)

func ClusterName(clusterName string) NacosOption {
	return func(r *nacosOptions) {
		r.clusterName = clusterName
	}
}

func GroupName(groupName string) NacosOption {
	return func(r *nacosOptions) {
		r.groupName = groupName
	}
}

func Weight(weight float64) NacosOption {
	return func(r *nacosOptions) {
		r.weight = weight
	}
}

func Enable(enable bool) NacosOption {
	return func(r *nacosOptions) {
		r.enable = enable
	}
}

func Healthy(healthy bool) NacosOption {
	return func(r *nacosOptions) {
		r.healthy = healthy
	}
}

func Ephemeral(ephemeral bool) NacosOption {
	return func(r *nacosOptions) {
		r.ephemeral = ephemeral
	}
}
