package context

import (
	"context"

	"google.golang.org/grpc"
)

func GRPCClientMiddleware(contextFunc func(ctx context.Context) context.Context) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req any, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if contextFunc != nil {
			ctx = contextFunc(ctx)
		}
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func GRPCServerMiddleware(contextFunc func(ctx context.Context) context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		if contextFunc != nil {
			ctx = contextFunc(ctx)
		}
		resp, err := handler(ctx, req)
		return resp, err
	}
}
