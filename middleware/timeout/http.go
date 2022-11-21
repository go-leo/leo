package timeout

import (
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

func GinMiddleware(duration time.Duration) gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(duration),
		timeout.WithHandler(func(c *gin.Context) { c.Next() }),
	)
}
