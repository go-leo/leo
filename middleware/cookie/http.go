package cookie

import (
	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx := r.Context()
		ctx = newContext(c, c)
		c.Request = r.WithContext(ctx)
	}
}
