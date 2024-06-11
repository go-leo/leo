package sdx

import (
	"context"
	"github.com/go-kit/kit/sd"
)

type InstancerFactory interface {
	New(ctx context.Context, endpointName string) sd.Instancer
}
