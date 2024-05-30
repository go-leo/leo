package endpointx

import (
	"context"
)

type nameKey struct{}

// InjectName injects the name into the context.
func InjectName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

// ExtractName extracts the name from the context.
func ExtractName(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}
