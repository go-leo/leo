package registry

import (
	"context"
	"net/url"
)

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
