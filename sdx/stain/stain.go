package stain

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type colorKey struct{}

// InjectColor injects the colors into the context.
func InjectColor(ctx context.Context, color string) context.Context {
	return context.WithValue(ctx, colorKey{}, color)
}

// ExtractColor extracts the colors from the context.
func ExtractColor(ctx context.Context) (string, bool) {
	color, ok := ctx.Value(colorKey{}).(string)
	return color, ok
}

func MatchColor(ctx context.Context, color string) bool {
	v, ok := ExtractColor(ctx)
	if !ok {
		return false
	}
	return v == color
}

func WithColor(color string, do endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return do(InjectColor(ctx, color), request)
	}
}
