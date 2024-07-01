package sdx

import (
	"context"
)

type colorKey struct{}

// InjectColor injects the colors into the context.
func InjectColor(ctx context.Context, colors string) context.Context {
	return context.WithValue(ctx, colorKey{}, colors)
}

// ExtractColor extracts the colors from the context.
func ExtractColor(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(colorKey{}).(string)
	return v, ok
}
