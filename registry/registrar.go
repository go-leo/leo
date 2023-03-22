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

type RegistrarFactory interface {
	Create(uri *url.URL) (Registrar, error)
}
