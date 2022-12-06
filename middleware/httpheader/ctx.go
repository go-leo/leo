package httpheader

import (
	"context"
	"net/http"
)

type key struct{}

func FromContext(ctx context.Context) (http.Header, bool) {
	val, ok := ctx.Value(key{}).(http.Header)
	return val, ok
}

func NewContext(ctx context.Context, h http.Header) context.Context {
	return context.WithValue(ctx, key{}, h)
}
