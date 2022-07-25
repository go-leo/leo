package hystrix

import (
	"context"

	"github.com/afex/hystrix-go/hystrix"
	"google.golang.org/grpc"
)

func GRPCClientMiddleware(conf hystrix.CommandConfig) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		m := hystrix.GetCircuitSettings()
		if _, ok := m[method]; !ok {
			hystrix.ConfigureCommand(method, conf)
		}
		return hystrix.Do(method, func() (err error) {
			return invoker(ctx, method, req, reply, cc, opts...)
		}, nil)
	}
}
