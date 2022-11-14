package cookie

import (
	"context"

	"github.com/gin-gonic/gin"
)

type key struct{}

func newContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, key{}, c)
}

func fromContext(ctx context.Context) (c *gin.Context, ok bool) {
	c, ok = ctx.Value(key{}).(*gin.Context)
	return
}

func SetCookie(ctx context.Context, name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	c, ok := fromContext(ctx)
	if !ok {
		return
	}
	c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}
