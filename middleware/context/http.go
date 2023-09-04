package context

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/hmldd/leo/runner/net/http/client"
)

func HTTPClientMiddleware(contextFunc func(ctx context.Context) context.Context) client.Interceptor {
	return func(ctx context.Context, req any, reply any, info *client.HTTPInfo, invoke client.Invoker) error {
		if contextFunc != nil {
			ctx = contextFunc(ctx)
			info.Request = info.Request.WithContext(ctx)
		}
		return invoke(ctx, req, reply, info)
	}
}

func GinMiddleware(contextFunc func(ctx context.Context) context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		if contextFunc != nil {
			ctx := contextFunc(c.Request.Context())
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
