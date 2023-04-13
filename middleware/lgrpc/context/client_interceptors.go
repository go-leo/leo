package context

import (
	"context"

	"google.golang.org/grpc"
)

func UnaryClientInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.UnaryClientInterceptor {
	if contextFunc == nil {
		contextFunc = func(ctx context.Context) context.Context { return ctx }
	}
	return func(
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = contextFunc(ctx)
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}

func StreamClientInterceptor(
	contextFunc func(ctx context.Context) context.Context,
) grpc.StreamClientInterceptor {
	if contextFunc == nil {
		contextFunc = func(ctx context.Context) context.Context { return ctx }
	}
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		ctx = contextFunc(ctx)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
