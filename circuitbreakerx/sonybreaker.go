package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/sony/gobreaker"
)

func GoBreaker(factory func(endpointName string) (*gobreaker.CircuitBreaker, error)) endpoint.Middleware {
	group := lazyloadx.Group[*gobreaker.CircuitBreaker]{
		New: factory,
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			endpointName, ok := endpointx.NameExtractor(ctx)
			if !ok {
				return next(ctx, request)
			}
			cb, err, _ := group.Load(endpointName)
			if err != nil {
				return nil, statusx.Canceled(statusx.Message(errLoadBreaker, endpointName))
			}
			response, err := circuitbreaker.Gobreaker(cb)(next)(ctx, request)
			if errors.Is(err, gobreaker.ErrTooManyRequests) || errors.Is(err, gobreaker.ErrOpenState) {
				return nil, ErrCircuitOpen
			}
			return response, err
		}
	}
}
