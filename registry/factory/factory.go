package factory

import (
	"errors"
	"net/url"

	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/registry"
	"github.com/go-leo/leo/registry/consul"
	"github.com/go-leo/leo/registry/nacos"
)

func NewRegistrar(uri *url.URL) (registry.Registrar, error) {
	switch uri.Scheme {
	case consul.Scheme:
		factory := consul.RegistrarFactory{URI: uri}
		return factory.Create()
	case nacos.Scheme:
		factory := nacos.RegistrarFactory{URI: uri}
		return factory.Create()
	default:
		return nil, errors.New("not support this scheme " + uri.Scheme)
	}
}

func NewDiscovery(uri *url.URL) (registry.Discovery, error) {
	switch uri.Scheme {
	case consul.Scheme:
		factory := consul.DiscoveryFactory{URI: uri, Logger: log.Discard{}}
		return factory.Create()
	case nacos.Scheme:
		factory := nacos.DiscoveryFactory{URI: uri, Logger: log.Discard{}}
		return factory.Create()
	default:
		return nil, errors.New("not support this scheme " + uri.Scheme)
	}
}
