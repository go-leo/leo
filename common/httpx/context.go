package httpx

import (
	"context"
	"net/http"
)

// Deprecated: Do not use. use github.com/go-leo/netx instead.
type clientKey struct{}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func ClientFromContext(ctx context.Context) (*http.Client, bool) {
	cli, ok := ctx.Value(clientKey{}).(*http.Client)
	return cli, ok
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func NewContextWithClient(ctx context.Context, cli *http.Client) context.Context {
	return context.WithValue(ctx, clientKey{}, cli)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
type requestKey struct{}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func RequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value(requestKey{}).(*http.Request)
	return req, ok
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func NewContextWithRequest(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, requestKey{}, req)
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
type responseKey struct{}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func ResponseFromContext(ctx context.Context) (*http.Response, bool) {
	resp, ok := ctx.Value(responseKey{}).(*http.Response)
	return resp, ok
}

// Deprecated: Do not use. use github.com/go-leo/netx instead.
func NewContextWithResponse(ctx context.Context, resp *http.Response) context.Context {
	return context.WithValue(ctx, responseKey{}, resp)
}
