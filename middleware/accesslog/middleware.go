package accesslog

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-leo/leo/v3/endpointx"
)

func Middleware(limiters map[string]ratelimit.Waiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			_ = endpointName
			// continue handle the request
			return next(ctx, request)
		}
	}
}
