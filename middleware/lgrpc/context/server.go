package context

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.UnaryServerInterceptor {
	if contextFunc == nil {
		contextFunc = func(ctx context.Context) context.Context { return ctx }
	}
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		ctx = contextFunc(ctx)
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func StreamServerInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.StreamServerInterceptor {
	if contextFunc == nil {
		contextFunc = func(ctx context.Context) context.Context { return ctx }
	}
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := stream.Context()
		ctx = contextFunc(ctx)
		return handler(srv, stream)
	}
}
