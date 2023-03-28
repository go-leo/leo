package registry

import (
	"context"
)

// Discovery is service Discovery
type Discovery interface {
	// GetInstances gets all service instances associated with a particular service name.
	GetInstances(ctx context.Context, instance ServiceInstance) ([]ServiceInstance, error)
	// GetServices return all known service names.
	GetServices(ctx context.Context) ([]string, error)
	// Watch whe service instances associated with a particular service name.
	Watch(ctx context.Context, instance ServiceInstance) (Watcher, error)
}

type Watcher interface {
	Notify(instanceC chan<- []ServiceInstance)
	StopNotify(instanceC chan<- []ServiceInstance)
	Close(ctx context.Context) error
}

type DiscoveryFactory interface {
	Create() (Discovery, error)
}

type DiscoveryFactoryFunc func() (Discovery, error)

func (f DiscoveryFactoryFunc) Create() (Discovery, error) {
	return f()
}
