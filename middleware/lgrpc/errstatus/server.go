package grpcerr

import (
	"context"

	"google.golang.org/grpc"
)

func ServerInterceptor(opts ...Option) (grpc.UnaryServerInterceptor, grpc.StreamServerInterceptor) {
	o := defaultOptions()
	o.apply(opts...)
	return unaryServerInterceptor(o),
		streamServerInterceptor(o)
}

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := defaultOptions()
	o.apply(opts...)
	return unaryServerInterceptor(o)
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := defaultOptions()
	o.apply(opts...)
	return streamServerInterceptor(o)
}

func unaryServerInterceptor(o *options) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		status := o.errorFunc(err)
		if status != nil {
			return nil, status.Err()
		}
		return resp, nil
	}
}

func streamServerInterceptor(o *options) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := handler(srv, stream)
		status := o.errorFunc(err)
		if status != nil {
			return status.Err()
		}
		return nil
	}
}
