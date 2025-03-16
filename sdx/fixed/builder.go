package fixed

import (
	"context"
	"github.com/go-kit/kit/sd"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net"
	"net/url"
)

var _ sdx.Builder = (*Builder)(nil)

const schemeName = "fixed"

type Builder struct{}

func (Builder) Scheme() string {
	return schemeName
}

func (Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string, logger kitlog.Logger) (sd.Instancer, error) {
	return sd.FixedInstancer{internal.ExtractEndpoint(instance)}, nil
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string, logger kitlog.Logger) (sd.Registrar, error) {
	return nil, nil
}
