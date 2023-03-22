package registry

import (
	"context"
	"errors"
	"net/url"
)

// ErrServiceNotFound not found service.
var ErrServiceNotFound = errors.New("service not found")

// Discovery is service Discovery
type Discovery interface {
	// GetInstances gets all service instances associated with a particular service name.
	GetInstances(ctx context.Context, serviceName string) ([]ServiceInstance, error)
	// GetServices return all known service names.
	GetServices(ctx context.Context) ([]string, error)
	// Watch whe service instances associated with a particular service name.
	Watch(ctx context.Context, serviceName string) (<-chan []ServiceInstance, error)
	// Stop watch service instances
	Stop(ctx context.Context, serviceName string, watcher <-chan []ServiceInstance) error
}

type DiscoveryFactory interface {
	Create(uri *url.URL) (Discovery, error)
}

type DiscoveryFactoryFunc func(uri *url.URL) (Discovery, error)

func (f DiscoveryFactoryFunc) Create(uri *url.URL) (Discovery, error) {
	return f(uri)
}
