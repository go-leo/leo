package trace

import (
	"context"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	traceInterceptor := otelgrpc.UnaryClientInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
	)
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		for _, skip := range o.Skips {
			if skip(ctx, method) {
				return invoker(ctx, method, req, reply, cc, opts...)
			}
		}
		return traceInterceptor(ctx, method, req, reply, cc, invoker, opts...)
	}
}

func StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	traceInterceptor := otelgrpc.StreamClientInterceptor(
		otelgrpc.WithPropagators(o.Propagators),
		otelgrpc.WithTracerProvider(o.TracerProvider),
	)
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		for _, skip := range o.Skips {
			if skip(ctx, method) {
				return streamer(ctx, desc, cc, method, opts...)
			}
		}
		return traceInterceptor(ctx, desc, cc, method, streamer)
	}
}
