package factory

import (
	"errors"
	"net/url"

	"github.com/go-leo/leo/common/consulx"
	"github.com/go-leo/leo/common/nacosx"
	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/registry"
	"github.com/go-leo/leo/registry/consul"
	"github.com/go-leo/leo/registry/nacos"
)

func NewRegistrar(uri *url.URL) (registry.Registrar, error) {
	switch uri.Scheme {
	case consul.Scheme:
		client, err := consulx.NewClient(uri)
		if err != nil {
			return nil, err
		}
		return consul.NewRegistrar(client), nil
	case nacos.Scheme:
		client, err := nacosx.NewNacosNamingClient(uri)
		if err != nil {
			return nil, err
		}
		return nacos.NewRegistrar(client), nil
	default:
		return nil, errors.New("not support this scheme " + uri.Scheme)
	}
}

func NewDiscovery(uri *url.URL) (registry.Discovery, error) {
	switch uri.Scheme {
	case consul.Scheme:
		client, err := consulx.NewClient(uri)
		if err != nil {
			return nil, err
		}
		return consul.NewDiscovery(client, log.Discard{}), nil
	case nacos.Scheme:
		client, err := nacosx.NewNacosNamingClient(uri)
		if err != nil {
			return nil, err
		}
		return nacos.NewDiscovery(client, log.Discard{}), nil
	default:
		return nil, errors.New("not support this scheme " + uri.Scheme)
	}
}
