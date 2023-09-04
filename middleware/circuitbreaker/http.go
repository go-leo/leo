package circuitbreaker

import (
	"context"
	"errors"

	"github.com/streadway/handy/breaker"

	"github.com/hmldd/leo/common/syncx"
	"github.com/hmldd/leo/runner/net/http/client"
)

var ErrCircuitOpen = errors.New("circuit open")

func HTTPClientMiddleware(opts ...Option) client.Interceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	breakers, _ := syncx.NewOnceCreator[Breaker](func() Breaker { return o.breakerCreator() })
	return func(ctx context.Context, req any, reply any, info *client.HTTPInfo, invoke client.Invoker) error {
		key := info.Path
		cb := breakers.LoadOrCreate(key)
		if !cb.Allow() {
			return breaker.ErrCircuitOpen
		}
		err := invoke(ctx, req, reply, info)
		if err == nil {
			cb.Success()
		} else {
			cb.Failure()
		}
		return err
	}
}
