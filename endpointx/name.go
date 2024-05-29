package endpointx

import "context"

type nameKey struct{}

func NewContext(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

func FromContext(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}
