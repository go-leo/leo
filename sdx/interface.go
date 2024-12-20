package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"net/url"
)

type InstancerFactory interface {
	// New news an instancer for the given target and color.
	// target represents a target for rpc, as specified in:
	// https://github.com/grpc/grpc/blob/master/doc/naming.md.
	New(ctx context.Context, target *url.URL, color string) (sd.Instancer, error)

	// Scheme returns the scheme supported by this resolver.  Scheme is defined
	// at https://github.com/grpc/grpc/blob/master/doc/naming.md.  The returned
	// string should not contain uppercase characters, as they will not match
	// the parsed target's scheme as defined in RFC 3986.
	Scheme() string
}
