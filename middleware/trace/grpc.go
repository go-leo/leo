package trace

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func GRPCClientMiddleware(opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	return otelgrpc.UnaryClientInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
		otelgrpc.WithInterceptorFilter(func(info *otelgrpc.InterceptorInfo) bool {
			if _, ok := o.Skips[info.Method]; ok {
				return false
			}
			return true
		}),
	)
}

func GRPCServerMiddleware(opts ...Option) grpc.UnaryServerInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	return otelgrpc.UnaryServerInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
		otelgrpc.WithInterceptorFilter(func(info *otelgrpc.InterceptorInfo) bool {
			if _, ok := o.Skips[info.UnaryServerInfo.FullMethod]; ok {
				return false
			}
			return true
		}),
	)
}
