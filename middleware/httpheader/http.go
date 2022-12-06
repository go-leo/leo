package httpheader

import (
	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(NewContext(c.Request.Context(), c.Request.Header))
	}
}
