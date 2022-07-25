package recovery

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinMiddleware(handlers ...func(*gin.Context, any)) gin.HandlerFunc {
	var handle func(*gin.Context, any)
	if len(handlers) == 0 {
		handle = func(c *gin.Context, err any) {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	} else {
		handle = handlers[0]
	}
	return gin.CustomRecoveryWithWriter(io.Discard, handle)
}
