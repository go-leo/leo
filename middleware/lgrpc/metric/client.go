package metric

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"google.golang.org/grpc"
	"time"
)

func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	clientStartedCounter, e := o.MeterProvider.Meter(kInstrumentationName).
		Int64Counter("grpc.client.started",
			metric.WithDescription("Total number of RPCs started on the client."))
	if e != nil {
		otel.Handle(e)
		panic(e)
	}

	clientHandledCounter, e := o.MeterProvider.Meter(kInstrumentationName).
		Int64Counter("grpc.client.handled",
			metric.WithDescription("Total number of RPCs completed by the client, regardless of success or failure."))
	if e != nil {
		otel.Handle(e)
		panic(e)
	}

	clientStreamMsgReceived, e := o.MeterProvider.Meter(kInstrumentationName).
		Int64Counter("grpc.client.msg.received",
			metric.WithDescription("Total number of RPC stream messages received by the client."))
	if e != nil {
		otel.Handle(e)
		panic(e)
	}

	clientStreamMsgSent, e := o.MeterProvider.Meter(kInstrumentationName).
		Int64Counter("grpc.client.msg.sent",
			metric.WithDescription("Total number of gRPC stream messages sent by the client."))
	if e != nil {
		otel.Handle(e)
		panic(e)
	}

	clientHandledHistogramOpts := []metric.Float64HistogramOption{
		metric.WithDescription("Histogram of response latency (seconds) of the gRPC until it is finished by the application."),
		metric.WithUnit("ms"),
	}
	if len(o.BucketBoundaries) > 0 {
		clientHandledHistogramOpts = append(clientHandledHistogramOpts, metric.WithExplicitBucketBoundaries(o.BucketBoundaries...))
	}
	clientHandledHistogram, e := o.MeterProvider.Meter(kInstrumentationName).
		Float64Histogram("grpc.client.handling.seconds",
			clientHandledHistogramOpts...)
	if e != nil {
		otel.Handle(e)
		panic(e)
	}

	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if _, ok := skipMap[method]; ok {
			return invoker(ctx, method, req, reply, cc, opts...)
		}
		// 开始时间
		startTime := time.Now()

		// 包含接口信息的属性
		attrs := append([]attribute.KeyValue{gRPCType()}, parseFullMethod(method)...)
		opt := metric.WithAttributes(attrs...)

		clientStartedCounter.Add(ctx, 1, opt)
		clientStreamMsgSent.Add(ctx, 1, opt)

		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			clientStreamMsgReceived.Add(ctx, 1, opt)
		}
		// 请求延迟直方图记录延迟
		elapsedTime := float64(time.Since(startTime)) / float64(time.Millisecond)
		clientHandledHistogram.Record(ctx, elapsedTime, opt)
		// 请求计数器加1
		clientHandledCounter.Add(ctx, 1, metric.WithAttributes(append(attrs, parseError(err))...))
		return err

	}
}
