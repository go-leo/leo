package requestid

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientInterceptor(opts ...Option) grpc.UnaryClientInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		var requestID string

		// 1. from context
		requestID, _ = FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			return clientInvoker(ctx, method, req, reply, cc, invoker, opts, o, requestID)
		}

		// 2. from grpc incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			return clientInvoker(NewContext(ctx, requestID), method, req, reply, cc, invoker, opts, o, requestID)
		}

		// 3. from TraceContext, TraceContext is a propagator that supports the W3C Trace Context format
		// (https://www.w3.org/TR/trace-context/)
		requestID, _ = fromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return clientInvoker(NewContext(ctx, requestID), method, req, reply, cc, invoker, opts, o, requestID)
		}

		// 4. generate
		requestID = o.generator()
		return clientInvoker(NewContext(ctx, requestID), method, req, reply, cc, invoker, opts, o, requestID)
	}
}

func clientInvoker(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts []grpc.CallOption, o *options, requestID string) error {
	ctx = toOutgoing(ctx, o, requestID)
	o.handler(ctx, requestID)
	return invoker(ctx, method, req, reply, cc, opts...)
}

func toOutgoing(ctx context.Context, o *options, v string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, o.headerKey, v)
}

func StreamClientInterceptor(opts ...Option) grpc.StreamClientInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		var requestID string

		// 1. from context
		requestID, _ = FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			return clientStreamer(ctx, desc, cc, method, streamer, opts, o, requestID)
		}

		// 2. from grpc incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			ctx = NewContext(ctx, requestID)
			return clientStreamer(ctx, desc, cc, method, streamer, opts, o, requestID)
		}

		// 3. from TraceContext, TraceContext is a propagator that supports the W3C Trace Context format
		// (https://www.w3.org/TR/trace-context/)
		requestID, _ = fromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			ctx = NewContext(ctx, requestID)
			return clientStreamer(ctx, desc, cc, method, streamer, opts, o, requestID)
		}

		// 4. generate
		requestID = o.generator()
		ctx = NewContext(ctx, requestID)
		return clientStreamer(ctx, desc, cc, method, streamer, opts, o, requestID)
	}
}

func clientStreamer(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts []grpc.CallOption, o *options, requestID string) (grpc.ClientStream, error) {
	ctx = toOutgoing(ctx, o, requestID)
	o.handler(ctx, requestID)
	return streamer(ctx, desc, cc, method, opts...)
}
