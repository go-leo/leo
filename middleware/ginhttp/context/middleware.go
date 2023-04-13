package context

import (
	"github.com/gin-gonic/gin"
)

func Middleware(opts ...Option) gin.HandlerFunc {
	o := defaultOptions()
	o.apply(opts...)
	return func(c *gin.Context) {
		ctx := o.contextFunc(c.Request.Context())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
