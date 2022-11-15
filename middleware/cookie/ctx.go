package cookie

import (
	"context"
	"errors"

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

func GetCookie(ctx context.Context, name string) (string, error) {
	c, ok := fromContext(ctx)
	if !ok {
		return "", errors.New("not found gin context")
	}
	return c.Cookie(name)
}

func SetCookie(ctx context.Context, name, value string, maxAge int, path, domain string, secure, httpOnly bool) error {
	c, ok := fromContext(ctx)
	if !ok {
		return errors.New("not found gin context")
	}
	c.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
	return nil
}
