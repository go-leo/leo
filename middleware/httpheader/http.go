package httpheader

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/runner/net/http/header"
)

func GinMiddleware() gin.HandlerFunc {
	return header.GinMiddleware()
}
