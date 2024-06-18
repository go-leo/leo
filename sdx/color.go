package sdx

import (
	"context"
)

type Color string

type colorKey struct{}

// InjectColor injects the colors into the context.
func InjectColor(ctx context.Context, colors Color) context.Context {
	return context.WithValue(ctx, colorKey{}, colors)
}

// ExtractColor extracts the colors from the context.
func ExtractColor(ctx context.Context) (Color, bool) {
	v, ok := ctx.Value(colorKey{}).(Color)
	return v, ok
}
