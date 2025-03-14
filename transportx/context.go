package transportx

import (
	"context"
)

type nameKey struct{}

// NameInjector injects the name into the context.
func NameInjector(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, nameKey{}, name)
}

// NameExtractor extracts the name from the context.
func NameExtractor(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(nameKey{}).(string)
	return v, ok
}
