package requestid

import (
	"context"

	"github.com/go-leo/gox/stringx"
	"google.golang.org/grpc"

	"codeup.aliyun.com/qimao/leo/leo/pkg/requestid"
)

func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		var requestID string
		// 1. from context
		requestID, _ = requestid.FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			o.handler(ctx, requestID)
			return handler(ctx, req)
		}
		// 2. from incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			ctx := requestid.NewContext(ctx, requestID)
			o.handler(ctx, requestID)
			return handler(ctx, req)
		}
		// // 3. from trace system traceID
		// requestID, _ = FromTrace(ctx)
		// if stringx.IsNotBlank(requestID) {
		// 	return NewContext(ctx, requestID)
		// }
		// 4. generate
		requestID = o.generator()
		ctx = requestid.NewContext(ctx, requestID)
		o.handler(ctx, requestID)
		return handler(ctx, req)
	}
}

func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := &options{}
	o.apply(opts...)
	o.init()
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
		var requestID string
		// 1. from context
		requestID, _ = requestid.FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			o.handler(ctx, requestID)
			return handler(srv, ss)
		}
		// 2. from incoming
		requestID, _ = fromIncoming(ctx, o)
		if stringx.IsNotBlank(requestID) {
			o.handler(ctx, requestID)
			return handler(srv, &serverStream{ctx: requestid.NewContext(ctx, requestID)})
		}
		// // 3. from trace system traceID
		// requestID, _ = FromTrace(ctx)
		// if stringx.IsNotBlank(requestID) {
		// 	return NewContext(ctx, requestID)
		// }
		// 4. generate
		requestID = o.generator()
		o.handler(ctx, requestID)
		return handler(srv, &serverStream{ctx: requestid.NewContext(ctx, requestID)})
	}
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *serverStream) Context() context.Context {
	return w.ctx
}
