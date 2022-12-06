package context

import (
	"context"

	"github.com/gin-gonic/gin"
)

func GinMiddleware(contextFunc func(ctx context.Context) context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		if contextFunc != nil {
			ctx := contextFunc(c.Request.Context())
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}
