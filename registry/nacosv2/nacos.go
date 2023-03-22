package nacosv2

import "github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"

type NamingClientFactory interface {
	Create() naming_client.INamingClient
}
