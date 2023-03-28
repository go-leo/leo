package nacosv2

import "github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"

type ConfigClientFactory interface {
	Create() (config_client.IConfigClient, error)
}

type ConfigClientFactoryFunc func() (config_client.IConfigClient, error)

func (f ConfigClientFactoryFunc) Create() (config_client.IConfigClient, error) {
	return f()
}
