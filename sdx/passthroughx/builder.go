package passthroughx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-leo/leo/v3/sdx"
)

const schemeName = "passthrough"

type InstancerBuilder struct{}

func (b *InstancerBuilder) Build(ctx context.Context, target *sdx.Target, color *sdx.Color) (sd.Instancer, error) {
	return NewInstancer(target.Instance()), nil
}

func (b *InstancerBuilder) Scheme() string {
	return schemeName
}

func NewInstancerBuilder() sdx.InstancerBuilder {
	return &InstancerBuilder{}
}
