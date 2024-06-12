package transportx

import "context"

type colorKey struct{}

type Color struct {
	Service string
	Colors  []string
}

// InjectColors injects the color into the context.
func InjectColors(ctx context.Context, colors []*Color) context.Context {
	return context.WithValue(ctx, colorKey{}, colors)
}

// ExtractColors extracts the color from the context.
func ExtractColors(ctx context.Context) ([]*Color, bool) {
	v, ok := ctx.Value(colorKey{}).([]*Color)
	return v, ok
}
