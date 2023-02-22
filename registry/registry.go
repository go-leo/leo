package registry

import (
	"context"
	"net/url"
)

// Service is Service
type Service interface {
	// ID is service id.
	ID() string
	// Name is service name.
	Name() string
	// Kind is service kind.
	Kind() string
	// Host is the host that Endpoint used.
	Host() string
	// Port is the port that Endpoint used.
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
	// Register the Service to registry.
	Register(ctx context.Context, service Service) error
	// Deregister the Service from registry.
	Deregister(ctx context.Context, service Service) error
}

// Discovery is service Discovery
type Discovery interface {
	// Scheme returns the scheme supported by this resolver.
	// Scheme is defined at https://github.com/grpc/grpc/blob/master/doc/naming.md.
	Scheme() string
	// GetService get service from service registry
	GetService(ctx context.Context, service Service) ([]Service, error)
	StartWatch(ctx context.Context, service Service) (<-chan []Service, error)
	StopWatch(ctx context.Context, service Service) error
}

type RegistrarFactory interface {
	Create(uri *url.URL) (Registrar, error)
}

type DiscoveryFactory interface {
	Create(uri *url.URL) (Discovery, error)
}
