package clientip

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/netx/addrx"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		r := c.Request
		ip := addrx.ClientPublicIP(r)
		if ip == "" {
			ip = addrx.ClientIP(r)
		}
		ctx := r.Context()
		ctx = NewContext(ctx, ip)
		c.Request = r.WithContext(ctx)
		c.Next()
	}
}

type key struct{}

func NewContext(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, key{}, ip)
}

func FromContext(ctx context.Context) (ip string, ok bool) {
	ip, ok = ctx.Value(key{}).(string)
	return
}
