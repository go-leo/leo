package registry

import "context"

type optionKey struct{}

func ContextWithParams(ctx context.Context, params map[string]any) context.Context {
	return context.WithValue(ctx, optionKey{}, params)
}

func ParamsFromContext(ctx context.Context) (map[string]any, bool) {
	v, ok := ctx.Value(optionKey{}).(map[string]any)
	return v, ok
}
