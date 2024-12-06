package configs

import (
	"context"
	"github.com/go-leo/leo/v3/config"
)

func Load(ctx context.Context, opts ...config.Option) (*Application, error) {
	return config.Load[*Application](ctx, opts...)
}

func Watch(ctx context.Context, ctxFunc func(ctx2 context.Context) context.Context, opts ...config.Option) (<-chan *Application, <-chan error, func(ctx context.Context)) {

	return out, errC, stopFunc
}
