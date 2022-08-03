package httpx

import (
	"context"
	"net/http"
)

type clientKey struct{}

func ClientFromContext(ctx context.Context) (*http.Client, bool) {
	cli, ok := ctx.Value(clientKey{}).(*http.Client)
	return cli, ok
}

func NewContextWithClient(ctx context.Context, cli *http.Client) context.Context {
	return context.WithValue(ctx, clientKey{}, cli)
}

type requestKey struct{}

func RequestFromContext(ctx context.Context) (*http.Request, bool) {
	req, ok := ctx.Value(requestKey{}).(*http.Request)
	return req, ok
}

func NewContextWithRequest(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, requestKey{}, req)
}

type responseKey struct{}

func ResponseFromContext(ctx context.Context) (*http.Response, bool) {
	resp, ok := ctx.Value(responseKey{}).(*http.Response)
	return resp, ok
}

func NewContextWithResponse(ctx context.Context, resp *http.Response) context.Context {
	return context.WithValue(ctx, responseKey{}, resp)
}
