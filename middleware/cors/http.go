package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GinMiddleware(config cors.Config) gin.HandlerFunc {
	return cors.New(config)
}
