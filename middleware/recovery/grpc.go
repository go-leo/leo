package recovery

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCServerMiddleware(handles ...func(context.Context, any) error) grpc.UnaryServerInterceptor {
	var handle func(context.Context, any) error
	if len(handles) == 0 {
		handle = func(ctx context.Context, p any) (err error) {
			return status.Errorf(codes.Internal, "panic triggered: %+v", err)
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
