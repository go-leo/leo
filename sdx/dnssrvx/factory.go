package dnssrvx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/dnssrv"
	"github.com/go-kit/log"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net/url"
	"time"
)

// schemeName for the urls
// All target URLs like 'dns:///instance' will be resolved by this builder
const schemeName = "dns"

type Factory struct {
	TTL    time.Duration
	Logger log.Logger
}

func (Factory) Scheme() string {
	return schemeName
}

func (f Factory) New(ctx context.Context, target *url.URL, color string) (sd.Instancer, error) {
	return dnssrv.NewInstancer(internal.ExtractEndpoint(target), f.TTL, f.Logger), nil
}
