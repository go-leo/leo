package registry

import "context"

type optionKey struct{}

func ContextWithParams(ctx context.Context, params map[string]any) context.Context {
	if params == nil {
		params = map[string]any{}
	}
	return context.WithValue(ctx, optionKey{}, params)
}

func ParamsFromContext(ctx context.Context) map[string]any {
	v, ok := ctx.Value(optionKey{}).(map[string]any)
	if ok {
		return v
	}
	return map[string]any{}
}
