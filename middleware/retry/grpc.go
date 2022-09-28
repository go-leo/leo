package circuitbreaker

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-leo/backoffx"

	"github.com/go-leo/retryx"
)

func GRPCClientMiddleware(maxAttempts uint, backoffFunc backoffx.BackoffFunc, isRetriable func(error) bool) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		return retryx.Call(ctx, maxAttempts, backoffFunc, func(_ int) error {
			err := invoker(ctx, method, req, reply, cc, opts...)
			if err != nil && isRetriable(err) {
				return err
			}
			return nil
		})
	}
}
