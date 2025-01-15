package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"net"
	"net/url"
)

// Builder is the interface that must be implemented by a service discovery.
// see: https://github.com/grpc/grpc/blob/master/doc/naming.md.
type Builder interface {
	// Scheme returns the scheme supported by this resolver.
	Scheme() string

	// BuildInstancer builds a sd.Instancer for the given instance information and color.
	BuildInstancer(ctx context.Context, instance *url.URL, color string) (sd.Instancer, error)

	// BuildRegistrar builds a sd.Registrar for the given instance information, address and color.
	BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string) (sd.Registrar, error)
}
