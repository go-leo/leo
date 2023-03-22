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
