package httpmethod

import (
	"context"
)

type key struct{}

func NewContext(ctx context.Context, method string) context.Context {
	return context.WithValue(ctx, key{}, method)
}

func FromContext(ctx context.Context) (method string, ok bool) {
	method, ok = ctx.Value(key{}).(string)
	return
}
