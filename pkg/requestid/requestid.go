package requestid

import "context"

type key struct{}

func FromContext(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(key{}).(string)
	return val, ok
}

func NewContext(ctx context.Context, v string) context.Context {
	return context.WithValue(ctx, key{}, v)
}
