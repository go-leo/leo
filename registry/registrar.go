package registry

import (
	"context"
)

// Registrar is service Registrar
type Registrar interface {
	// Register the ServiceInstance to registry.
	Register(ctx context.Context, instance ServiceInstance) error
	// Deregister the ServiceInstance from registry.
	Deregister(ctx context.Context, instance ServiceInstance) error
}
