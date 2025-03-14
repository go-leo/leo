package dnssrvx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/dnssrv"
	"github.com/go-kit/log"
	kitlog "github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net"
	"net/url"
	"time"
)

var _ sdx.Builder = (*Builder)(nil)

// schemeName for the urls
// All target URLs like 'dns:///instance' will be resolved by this builder
const schemeName = "dns"

type Builder struct {
	TTL    time.Duration
	Logger log.Logger
}

func (Builder) Scheme() string {
	return schemeName
}

func (b Builder) BuildInstancer(ctx context.Context, instance *url.URL, color string, logger kitlog.Logger) (sd.Instancer, error) {
	return dnssrv.NewInstancer(internal.ExtractEndpoint(instance), b.TTL, b.Logger), nil
}

func (b Builder) BuildRegistrar(ctx context.Context, instance *url.URL, ip net.IP, port int, color string, logger kitlog.Logger) (sd.Registrar, error) {
	//TODO implement me
	panic("implement me")
}
