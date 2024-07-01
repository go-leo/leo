package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
)

// Address is Address
type Address interface {
	Host() string
	Port() int
}

// RegistrarBuilder is a builder that can build an Registrar for a given target and address.
type RegistrarBuilder interface {
	// Build builds an instancer for the given target and address.
	// target represents a target for rpc, as specified in:
	// https://github.com/grpc/grpc/blob/master/doc/naming.md.
	// address represents the host:port address for rpc.
	Build(ctx context.Context, target *Target, address Address) sd.Registrar

	// Scheme returns the scheme supported by this resolver.  Scheme is defined
	// at https://github.com/grpc/grpc/blob/master/doc/naming.md.  The returned
	// string should not contain uppercase characters, as they will not match
	// the parsed target's scheme as defined in RFC 3986.
	Scheme() string
}
