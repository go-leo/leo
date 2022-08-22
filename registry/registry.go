package registry

import (
	"context"
)

const (
	TransportHTTP  = "http"
	TransportHTTPS = "https"
	TransportGRPC  = "gRPC"
)

// Registrar registers service information to a service registry.
type Registrar interface {
	// Register the ServiceInfo to registry.
	Register(ctx context.Context, service *ServiceInfo) error
	// Deregister the ServiceInfo from registry.
	Deregister(ctx context.Context, service *ServiceInfo) error
}

// Discovery registers service information to a service registry.
type Discovery interface {
	// Scheme returns the scheme supported by this resolver.
	// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
	Scheme() string
	GetService(ctx context.Context, service *ServiceInfo) ([]*ServiceInfo, error)
	Watch(ctx context.Context, service *ServiceInfo) (<-chan []*ServiceInfo, error)
	StopWatch(ctx context.Context, service *ServiceInfo) error
}

type RegistrarFactory interface {
	Create() (Registrar, error)
}

type DiscoveryFactory interface {
	Create() (Discovery, error)
}
