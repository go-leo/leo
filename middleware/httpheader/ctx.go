package httpheader

import (
	"context"
	"net/http"

	"github.com/go-leo/leo/runner/net/http/header"
)

func NewContext(ctx context.Context, h http.Header) context.Context {
	return header.NewContext(ctx, h)
}

func FromContext(ctx context.Context) (h http.Header, ok bool) {
	return header.FromContext(ctx)
}
