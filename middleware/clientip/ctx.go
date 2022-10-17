package clientip

import (
	"context"
)

type key struct{}

func NewContext(ctx context.Context, ip string) context.Context {
	return context.WithValue(ctx, key{}, ip)
}

func FromContext(ctx context.Context) (ip string, ok bool) {
	ip, ok = ctx.Value(key{}).(string)
	return
}
