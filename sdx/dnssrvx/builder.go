package dnssrvx

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
	TTL time.Duration
}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target) (sd.Instancer, error) {
	return dnssrv.NewInstancer(target.Instance(), b.TTL, logx.FromContext(ctx)), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}
