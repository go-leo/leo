package circuitbreakerx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/statusx"
)

// Breaker execute the endpoint, if the circuit is open, it will return false, nil, nil
// if the circuit is closed, it will execute the endpoint and return true, response, err.
// See HystrixBreaker,
type Breaker interface {
	Execute(ctx context.Context, request any, e endpoint.Endpoint) (bool, any, error)
}

// Newer create the Breaker. See HystrixNewer
type Newer interface {
	New() Breaker
}

// Factory is the Breaker factory.
type Factory struct {
	Newer Newer
}

func Middleware(factory Factory) endpoint.Middleware {
	breaker := factory.Newer.New()
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			ok, resp, err := breaker.Execute(ctx, request, next)
			if !ok {
				return nil, statusx.ErrUnavailable.With(statusx.Message("circuit open"))
			}
			return resp, err
		}
	}
}
