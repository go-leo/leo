package httpmethod

import (
	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		method := c.Request.Method
		ctx := r.Context()
		ctx = NewContext(ctx, method)
		c.Request = r.WithContext(ctx)
	}
}
