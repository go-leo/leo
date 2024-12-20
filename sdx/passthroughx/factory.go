package passthroughx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net/url"
)

var _ sdx.InstancerFactory = (*Factory)(nil)

const schemeName = "passthrough"

type Factory struct{}

func (Factory) Scheme() string {
	return schemeName
}

func (Factory) New(ctx context.Context, target *url.URL, color string) (sd.Instancer, error) {
	return Instancer{Instance: internal.ExtractEndpoint(target)}, nil
}
