package noop

import (
	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
