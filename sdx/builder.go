package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"net/url"
	"strings"
)

// InstancerBuilder is a builder that can build an instancer for a given target.
type InstancerBuilder interface {

	// Build builds an instancer for the given target and color.
	// target represents a target for rpc, as specified in:
	// https://github.com/grpc/grpc/blob/master/doc/naming.md.
	Build(ctx context.Context, target *Target, color *Color) (sd.Instancer, error)

	// Scheme returns the scheme supported by this resolver.  Scheme is defined
	// at https://github.com/grpc/grpc/blob/master/doc/naming.md.  The returned
	// string should not contain uppercase characters, as they will not match
	// the parsed target's scheme as defined in RFC 3986.
	Scheme() string
}

// Target represents a target for rpc, as specified in:
// https://github.com/grpc/grpc/blob/master/doc/naming.md.
// FROM google.golang.org/grpc/resolver/resolver.go
type Target struct {
	// URL contains the parsed dial target with an optional default scheme added
	// to it if the original dial target contained no scheme or contained an
	// unregistered scheme. Any query params specified in the original dial
	// target can be accessed from here.
	URL url.URL
}

// Instance retrieves instance without leading "/" from either `URL.Path`
// or `URL.Opaque`. The latter is used when the former is empty.
func (t Target) Instance() string {
	endpoint := t.URL.Path
	if endpoint == "" {
		endpoint = t.URL.Opaque
	}
	return strings.TrimPrefix(endpoint, "/")
}

// String returns the canonical string representation of Target.
func (t Target) String() string {
	return t.URL.Scheme + "://" + t.URL.Host + "/" + t.Instance()
}

// ParseTarget uses RFC 3986 semantics to parse the given target into a
// resolver.Target struct containing url. Query params are stripped from the
// endpoint.
func ParseTarget(target string) (*Target, error) {
	u, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	return &Target{URL: *u}, nil
}
