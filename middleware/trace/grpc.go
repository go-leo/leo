package trace

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func GRPCClientMiddleware(opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	traceInterceptor := otelgrpc.UnaryClientInterceptor(otelgrpc.WithPropagators(o.Propagators), otelgrpc.WithTracerProvider(o.TracerProvider))
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if _, ok := skipMap[method]; ok {
			return invoker(ctx, method, req, reply, cc, opts...)
		}
		return traceInterceptor(ctx, method, req, reply, cc, invoker, opts...)
	}
}

func GRPCServerMiddleware(opts ...Option) grpc.UnaryServerInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	traceInterceptor := otelgrpc.UnaryServerInterceptor(otelgrpc.WithPropagators(o.Propagators), otelgrpc.WithTracerProvider(o.TracerProvider))
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		if _, ok := skipMap[info.FullMethod]; ok {
			return handler(ctx, req)
		}
		return traceInterceptor(ctx, req, info, handler)
	}
}
