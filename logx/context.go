package logx

import (
	"context"
)

type kvKey struct{}

// FetchKeyVals fetches key values from the context.
func FetchKeyVals(ctx context.Context) []interface{} {
	v, _ := ctx.Value(kvKey{}).([]interface{})
	return v
}

// InjectKeyVals injects key values into the context.
func InjectKeyVals(ctx context.Context, keyvals ...interface{}) context.Context {
	return context.WithValue(ctx, kvKey{}, append(FetchKeyVals(ctx), keyvals...))
}
