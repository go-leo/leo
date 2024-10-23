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
	"github.com/streadway/handy/breaker"
)

func HandyBreaker(factory func(endpointName string) breaker.Breaker) endpoint.Middleware {
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
			cb := value.(breaker.Breaker)
			response, err := circuitbreaker.HandyBreaker(cb)(next)(ctx, request)
			if errors.Is(err, breaker.ErrCircuitOpen) {
				return nil, statusx.ErrUnavailable.With(statusx.Wrap(ErrCircuitOpen))
			}
			return response, err
		}
	}
}
