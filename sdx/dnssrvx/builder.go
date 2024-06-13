package consulx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/dnssrv"
	"github.com/go-leo/leo/v3/logx"
	"github.com/go-leo/leo/v3/sdx"
	"time"
)

// schemeName for the urls
// All target URLs like 'dns:///instance' will be resolved by this builder
const schemeName = "dns"

type InstancerBuilder struct {
	ttl time.Duration
}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target, color *sdx.Color) (sd.Instancer, error) {
	return dnssrv.NewInstancer(target.Instance(), b.ttl, logx.FromContext(ctx)), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}

func NewInstancerBuilder(ttl time.Duration) sdx.InstancerBuilder {
	return &InstancerBuilder{ttl: ttl}
}
