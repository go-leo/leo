package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

func Middleware(contextFunc func(ctx context.Context) context.Context) gin.HandlerFunc {
	if contextFunc == nil {
		contextFunc = func(ctx context.Context) context.Context { return ctx }
	}
	return func(c *gin.Context) {
		ctx := contextFunc(c.Request.Context())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
