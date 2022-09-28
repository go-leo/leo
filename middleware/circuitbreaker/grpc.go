package circuitbreaker

import (
	"context"

	"github.com/streadway/handy/breaker"
	"google.golang.org/grpc"

	"github.com/go-leo/syncx"
)

func GRPCClientMiddleware(opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	breakers, _ := syncx.NewOnceCreator[Breaker](func() Breaker { return o.breakerCreator() })
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		key := method
		cb := breakers.LoadOrCreate(key)
		if !cb.Allow() {
			return breaker.ErrCircuitOpen
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			cb.Success()
		} else {
			cb.Failure()
		}
		return err
	}
}
