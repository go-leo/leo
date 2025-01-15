package passthroughx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net"
	"net/url"
)

var _ sdx.Builder = (*Builder)(nil)

const schemeName = "passthrough"

type Builder struct{}

func (Builder) Scheme() string {
	return schemeName
}

func (Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string) (sd.Instancer, error) {
	return Instancer{Instance: internal.ExtractEndpoint(instance)}, nil
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string) (sd.Registrar, error) {

	return nil, nil
}
