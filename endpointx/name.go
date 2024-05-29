package endpointx

import (
	"context"
)

type nameKey struct{}

func InjectName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

func ExtractName(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}
