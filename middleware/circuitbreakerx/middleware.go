package circuitbreakerx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/statusx"
)

// Breaker execute the endpoint.
// if the circuit is open, it will return false, nil, nil
// if the circuit is closed, it will execute the endpoint and return true, response, err.
type Breaker interface {
	Execute(ctx context.Context, request any, next endpoint.Endpoint) (any, error, bool)
}

func Middleware(breaker Breaker) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			resp, err, ok := breaker.Execute(ctx, request, next)
			if !ok {
				return nil, statusx.ErrUnavailable.With(statusx.Message("circuit breaker is open"))
			}
			return resp, err
		}
	}
}
