package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ctx := r.Context()
		ctx = NewContext(ctx, r.Header)
		c.Request = r.WithContext(ctx)
		c.Next()
	}
}

type key struct{}

func NewContext(ctx context.Context, h http.Header) context.Context {
	return context.WithValue(ctx, key{}, h)
}

func FromContext(ctx context.Context) (h http.Header, ok bool) {
	h, ok = ctx.Value(key{}).(http.Header)
	return
}
