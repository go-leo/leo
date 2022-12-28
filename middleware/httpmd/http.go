package httpmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Metadata struct {
	Method   string
	Header   http.Header
	ClientIP string
}

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		md := Metadata{
			Method:   c.Request.Method,
			Header:   c.Request.Header,
			ClientIP: c.ClientIP(),
		}
		c.Request = c.Request.WithContext(NewContext(c.Request.Context(), &md))
	}
}
