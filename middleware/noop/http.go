package noop

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/hmldd/leo/runner/net/http/client"
)

func HTTPClientMiddleware() client.Interceptor {
	return func(ctx context.Context, req any, reply any, info *client.HTTPInfo, invoke client.Invoker) error {
		return invoke(ctx, req, reply, info)
	}
}

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
