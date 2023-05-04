package requestid

import (
	"context"
	"encoding/hex"
	"math/rand"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/stringx"

	middlewarecontext "github.com/go-leo/leo/middleware/context"
	"github.com/go-leo/leo/runner/net/http/client"
)

func HTTPClientMiddleware() client.Interceptor {
	return middlewarecontext.HTTPClientMiddleware(func(ctx context.Context) context.Context {
		// 1. from context
		requestID, ok := FromContext(ctx)
		if !ok || stringx.IsNotBlank(requestID) {
			return ToHeader(ctx, requestID)
		}
		// 2. from incoming
		requestID, ok = FromHeader(ctx)
		if !ok || stringx.IsNotBlank(requestID) {
			return ToHeader(ctx, requestID)
		}
		// 3. from trace system traceID
		requestID, ok = FromTrace(ctx)
		if !ok || stringx.IsNotBlank(requestID) {
			return ToHeader(ctx, requestID)
		}
		// 4. generate
		randSource := randPool.Get().(*rand.Rand)
		defer randPool.Put(randSource)
		var tid [16]byte
		randSource.Read(tid[:])
		requestID = hex.EncodeToString(tid[:])
		return ToHeader(ctx, requestID)
	})
}

func GinMiddleware() gin.HandlerFunc {
	return middlewarecontext.GinMiddleware(func(ctx context.Context) context.Context {
		var requestID string
		// 1. from context
		requestID, _ = FromContext(ctx)
		if stringx.IsNotBlank(requestID) {
			return ctx
		}
		// 2. from incoming
		requestID, _ = FromHeader(ctx)
		if stringx.IsNotBlank(requestID) {
			return NewContext(ctx, requestID)
		}
		// 3. from trace system traceID
		requestID, _ = FromTrace(ctx)
		if stringx.IsNotBlank(requestID) {
			return NewContext(ctx, requestID)
		}
		// 4. generate
		randSource := randPool.Get().(*rand.Rand)
		defer randPool.Put(randSource)
		var tid [16]byte
		randSource.Read(tid[:])
		requestID = hex.EncodeToString(tid[:])
		return NewContext(ctx, requestID)
	})
}
