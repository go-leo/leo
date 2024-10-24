package sdx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type targetKey struct{}

// InjectTarget injects the target into the context.
func InjectTarget(ctx context.Context, target string) context.Context {
	return context.WithValue(ctx, targetKey{}, target)
}

// ExtractTarget extracts the target from the context.
func ExtractTarget(ctx context.Context) (string, bool) {
	v, ok := ctx.Value(targetKey{}).(string)
	return v, ok
}

func WithTarget(target string, do endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		return do(InjectTarget(ctx, target), request)
	}
}