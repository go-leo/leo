package requestid

import (
	"context"

	"google.golang.org/grpc"

	"github.com/go-leo/stringx"

	middlewarecontext "github.com/go-leo/leo/v2/middleware/context"
)

func GRPCClientMiddleware() grpc.UnaryClientInterceptor {
	return middlewarecontext.GRPCClientMiddleware(func(ctx context.Context) context.Context {
		// 1. from context
		requestID, _ := FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			return ToOutgoing(ctx, requestID)
		}
		// 2. from grpc incoming
		requestID, _ = FromIncoming(ctx)
		if stringx.IsNotBlank(requestID) {
			return ToOutgoing(NewContext(ctx, requestID), requestID)
		}
		// 3. from http header
		requestID, _ = FromHeader(ctx)
		if stringx.IsNotBlank(requestID) {
			return ToOutgoing(NewContext(ctx, requestID), requestID)
		}
		// 4. from trace system traceID
		requestID, _ = FromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return ToOutgoing(NewContext(ctx, requestID), requestID)
		}
		// 5. generate
		requestID = Generate()
		return ToOutgoing(NewContext(ctx, requestID), requestID)
	})
}

func GRPCServerMiddleware() grpc.UnaryServerInterceptor {
	return middlewarecontext.GRPCServerMiddleware(func(ctx context.Context) context.Context {
		var requestID string
		// 1. from context
		requestID, _ = FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			return ctx
		}
		// 2. from incoming
		requestID, _ = FromIncoming(ctx)
		if stringx.IsNotBlank(requestID) {
			return NewContext(ctx, requestID)
		}
		// 3. from trace system traceID
		requestID, _ = FromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return NewContext(ctx, requestID)
		}
		// 4. generate
		requestID = Generate()
		return NewContext(ctx, requestID)
	})
}
