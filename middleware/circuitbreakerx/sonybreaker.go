package circuitbreakerx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/sony/gobreaker"
)

func GoBreaker(factory func(endpointName string) *gobreaker.CircuitBreaker) endpoint.Middleware {
	group := lazyloadx.Group{
		New: func(key string) (any, error) {
			return factory(key), nil
		},
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			value, err, _ := group.Load(endpointName)
			if err != nil {
				panic(fmt.Errorf("circuitbreakerx: failed to load %s breaker", endpointName))
			}
			cb := value.(*gobreaker.CircuitBreaker)
			response, err := circuitbreaker.Gobreaker(cb)(next)(ctx, request)
			if errors.Is(err, gobreaker.ErrTooManyRequests) || errors.Is(err, gobreaker.ErrOpenState) {
				return nil, statusx.ErrUnavailable.With(statusx.Wrap(ErrCircuitOpen))
			}
			return response, err
		}
	}
}
