package passthroughx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
	"github.com/go-leo/leo/v3/sdx/internal"
	"net/url"
)

var _ sdx.Builder = (*Builder)(nil)

const schemeName = "passthrough"

type Builder struct{}

func (Builder) Scheme() string {
	return schemeName
}

func (Builder) BuildInstancer(ctx context.Context, target *url.URL, color string) (sd.Instancer, error) {
	return Instancer{Instance: internal.ExtractEndpoint(target)}, nil
}

func (b Builder) BuildRegistrar(ctx context.Context, target *url.URL, address sdx.Address, color string) (sd.Registrar, error) {

	return nil, nil
}
