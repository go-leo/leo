package envx

import (
	"context"
	"github.com/go-leo/leo/v3/configx"
)

var _ configx.Resource = (*Resource)(nil)

type Resource struct {
}

func (r *Resource) Load(ctx context.Context) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Resource) Format() string {
	return "env"
}

func (r *Resource) Watch(ctx context.Context, notifyC chan<- *configx.Event) (func(), error) {
	//TODO implement me
	panic("implement me")
}
