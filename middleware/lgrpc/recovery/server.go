package recovery

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/runtimex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerInterceptor(handles ...func(context.Context, any) error) grpc.UnaryServerInterceptor {
	var handle func(context.Context, any) error
	if len(handles) == 0 {
		handle = func(ctx context.Context, p any) error {
			return status.Errorf(codes.Internal, "panic triggered: %+v, stack: %s", p, runtimex.Stack(0))
		}
	} else {
		handle = handles[0]
	}
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = handle(ctx, r)
			}
		}()

		resp, err := handler(ctx, req)
		panicked = false
		return resp, err
	}
}

func StreamServerInterceptor(handles ...func(context.Context, any) error) grpc.StreamServerInterceptor {
	var handle func(context.Context, any) error
	if len(handles) == 0 {
		handle = func(ctx context.Context, p any) error {
			return status.Errorf(codes.Internal, "panic triggered: %+v, stack: %s", p, runtimex.Stack(0))
		}
	} else {
		handle = handles[0]
	}
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = handle(stream.Context(), r)
			}
		}()

		err = handler(srv, stream)
		panicked = false
		return err
	}
}
