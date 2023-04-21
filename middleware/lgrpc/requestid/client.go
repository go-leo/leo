package requestid

import (
	"context"

	"github.com/go-leo/gox/stringx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"codeup.aliyun.com/qimao/leo/leo/pkg/requestid"
)

// UnaryClientInterceptor creates the unary client interceptor wrapped with Sentinel entry.
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
		requestID, _ = requestid.FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			ctx = toOutgoing(ctx, o, requestID)
			o.handler(ctx, requestID)
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		// 2. from grpc incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
			o.handler(ctx, requestID)
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		// // 3. from trace system traceID
		// requestID, _ = fromTrace(ctx)
		// if stringx.IsNotBlank(requestID) {
		// 	ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
		// 	return invoker(ctx, method, req, reply, cc, opts...)
		// }

		// 4. generate
		requestID = o.generator()
		ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
		o.handler(ctx, requestID)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func fromIncoming(ctx context.Context, o *options) (string, bool) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		vals := md.Get(o.headerKey)
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return "", false
}

// func fromTrace(ctx context.Context) (string, bool) {
// 	spanContext := trace.SpanContextFromContext(ctx)
// 	if spanContext.HasTraceID() {
// 		return spanContext.TraceID().String(), true
// 	}
// 	return "", false
// }

func toOutgoing(ctx context.Context, o *options, v string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, o.headerKey, v)
}

// StreamClientInterceptor creates the stream client interceptor wrapped with Sentinel entry.
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
		requestID, _ = requestid.FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			ctx = toOutgoing(ctx, o, requestID)
			o.handler(ctx, requestID)
			return streamer(ctx, desc, cc, method, opts...)
		}

		// 2. from grpc incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
			o.handler(ctx, requestID)
			return streamer(ctx, desc, cc, method, opts...)
		}

		// // 3. from trace system traceID
		// requestID, _ = fromTrace(ctx)
		// if stringx.IsNotBlank(requestID) {
		// 	ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
		// 	return invoker(ctx, method, req, reply, cc, opts...)
		// }

		// 4. generate
		requestID = o.generator()
		ctx = toOutgoing(requestid.NewContext(ctx, requestID), o, requestID)
		o.handler(ctx, requestID)
		return streamer(ctx, desc, cc, method, opts...)
	}
}
