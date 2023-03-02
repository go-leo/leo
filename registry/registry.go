package registry

import (
	"context"
	"net/url"
)

// ServiceInstance represents an instance of a service in a discovery system.
type ServiceInstance interface {
	// InstanceID is the unique instance ID as registered.
	InstanceID() string
	// ID is the service ID as registered.
	ID() string
	// Name is service name.
	Name() string
	// Kind is service kind.
	Kind() string
	// Host is the hostname of the registered service instance.
	Host() string
	// Port is the port of the registered service instance.
	Port() int
	// Metadata is other information the service carried.
	Metadata() map[string]string
	// Version is version of the service
	Version() string
	// Address is host:port
	Address() string
}

// Registrar is service Registrar
type Registrar interface {
	// Register the ServiceInstance to registry.
	Register(ctx context.Context, service ServiceInstance) error
	// Deregister the ServiceInstance from registry.
	Deregister(ctx context.Context, service ServiceInstance) error
}

// Discovery is service Discovery
type Discovery interface {
	// Scheme returns the scheme supported by this resolver.
	// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
	Scheme() string
	// GetService get service from service registry
	GetService(ctx context.Context, service ServiceInstance) ([]ServiceInstance, error)
	StartWatch(ctx context.Context, service ServiceInstance) (<-chan []ServiceInstance, error)
	StopWatch(ctx context.Context, service ServiceInstance) error
}

type RegistrarFactory interface {
	Create(uri *url.URL) (Registrar, error)
}

type DiscoveryFactory interface {
	Create(uri *url.URL) (Discovery, error)
}
