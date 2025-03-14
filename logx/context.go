package logx

import (
	"context"
)

type kvKey struct{}

// FetchKeyValsExtractor fetches key values from the context.
func FetchKeyValsExtractor(ctx context.Context) []interface{} {
	v, _ := ctx.Value(kvKey{}).([]interface{})
	return v
}

// KeyValsExtractorInjector injects key values into the context.
func KeyValsExtractorInjector(ctx context.Context, keyvals ...interface{}) context.Context {
	return context.WithValue(ctx, kvKey{}, append(FetchKeyValsExtractor(ctx), keyvals...))
}
