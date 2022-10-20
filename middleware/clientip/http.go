package clientip

import (
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
