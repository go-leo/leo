package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"net/url"
)

// Address is Address
type Address interface {
	Host() string
	Port() string
}

// Builder is the interface that must be implemented by a service discovery.
// target represents a target for rpc, as specified in:
// https://github.com/grpc/grpc/blob/master/doc/naming.md.
type Builder interface {
	// Scheme returns the scheme supported by this resolver.
	Scheme() string

	// BuildInstancer builds an instancer for the given target and color.
	BuildInstancer(ctx context.Context, target *url.URL, color string) (sd.Instancer, error)

	// BuildRegistrar builds a Registrar for the given target, address and color.
	BuildRegistrar(ctx context.Context, target *url.URL, address Address, color string) (sd.Registrar, error)
}
