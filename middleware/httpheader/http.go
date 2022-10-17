package httpheader

import (
	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx := r.Context()
		ctx = NewContext(ctx, r.Header)
		c.Request = r.WithContext(ctx)
		c.Next()
	}
}
