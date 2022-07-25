package nacosx

import (
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// nacos://ip1:port1,ip2:port2?contentType=yaml&namespace=ns&group=g&dataId=d

func NewNacosConfigClient(uri *url.URL) (config_client.IConfigClient, error) {
	var sc []constant.ServerConfig
	hosts := strings.Split(uri.Host, ",")
	for _, host := range hosts {
		ipAddr, portStr, err := net.SplitHostPort(host)
		if err != nil {
			return nil, fmt.Errorf("failed split host port %s, %w", host, err)
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("failed convert port to int %s, %w", portStr, err)
		}
		sc = append(sc, *constant.NewServerConfig(ipAddr, uint64(port)))
	}

	namespace := uri.Query().Get("namespace")
	cc := constant.NewClientConfig(
		constant.WithNamespaceId(namespace),
		constant.WithNotLoadCacheAtStart(true),
	)

	clientParam := vo.NacosClientParam{
		ClientConfig:  cc,
		ServerConfigs: sc,
	}
	client, err := clients.NewConfigClient(clientParam)
	if err != nil {
		return nil, fmt.Errorf("failed new nacos client, %w", err)
	}
	return client, err
}

func NewNacosNamingClient(uri *url.URL) (naming_client.INamingClient, error) {
	var sc []constant.ServerConfig
	hosts := strings.Split(uri.Host, ",")
	for _, host := range hosts {
		ipAddr, portStr, err := net.SplitHostPort(host)
		if err != nil {
			return nil, fmt.Errorf("failed split host port %s, %w", host, err)
		}
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("failed convert port to int %s, %w", portStr, err)
		}
		sc = append(sc, *constant.NewServerConfig(ipAddr, uint64(port)))
	}

	cc := constant.NewClientConfig(constant.WithNotLoadCacheAtStart(true))

	clientParam := vo.NacosClientParam{ClientConfig: cc, ServerConfigs: sc}
	client, err := clients.NewNamingClient(clientParam)
	if err != nil {
		return nil, fmt.Errorf("failed new nacos client, %w", err)
	}
	return client, err
}
