package context

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := defaultOptions()
	o.apply(opts...)
	return unaryServerInterceptor(o.contextFunc)
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := defaultOptions()
	o.apply(opts...)
	return streamServerInterceptor(o.contextFunc)
}

func unaryServerInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.UnaryServerInterceptor {
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

func streamServerInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.StreamServerInterceptor {
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
