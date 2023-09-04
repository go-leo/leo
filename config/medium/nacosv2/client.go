package nacosv2

import (
	"os"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func NewClient(host string, port uint64, namespace string) (config_client.IConfigClient, error) {
	sc := make([]constant.ServerConfig, 1)
	sc[0] = *constant.NewServerConfig(host, port)

	ops := make([]constant.ClientOption, 0, 3)
	ops = append(
		ops,
		constant.WithNamespaceId(namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
	)
	ak := os.Getenv("NACOS_ACCESS_KEY")
	sk := os.Getenv("NACOS_SECRET_KEY")
	if ak != "" && sk != "" {
		ops = append(ops, constant.WithAccessKey(ak), constant.WithSecretKey(sk))
	}

	cc := constant.NewClientConfig(ops...)

	clientParam := vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
	client, err := clients.NewConfigClient(clientParam)
	return client, err
}
