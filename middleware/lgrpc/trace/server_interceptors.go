package trace

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	traceInterceptor := otelgrpc.UnaryServerInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
	)
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		for _, skip := range o.Skips {
			if skip(ctx, info.FullMethod) {
				return handler(ctx, req)
			}
		}
		return traceInterceptor(ctx, req, info, handler)
	}
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	traceInterceptor := otelgrpc.StreamServerInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
	)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		for _, skip := range o.Skips {
			if skip(ss.Context(), info.FullMethod) {
				return handler(srv, ss)
			}
		}
		return traceInterceptor(srv, ss, info, handler)
	}
}
