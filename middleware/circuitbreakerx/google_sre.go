package circuitbreakerx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
)

// GoogleSreBreaker
// see: github.com/go-kratos/aegis/circuitbreaker/sre
// see: https://landing.google.com/sre/sre-book/chapters/handling-overload/
func GoogleSreBreaker(factory func(endpointName string) circuitbreaker.CircuitBreaker) endpoint.Middleware {
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
			cb := value.(circuitbreaker.CircuitBreaker)
			if err := cb.Allow(); err != nil {
				cb.MarkFailed()
				return nil, statusx.ErrUnavailable.With(statusx.Wrap(ErrCircuitOpen))
			}
			response, err := next(ctx, request)
			if err != nil && (statusx.ErrInternal.Equals(err) || statusx.ErrUnavailable.Equals(err) || statusx.ErrDeadlineExceeded.Equals(err)) {
				cb.MarkFailed()
			} else {
				cb.MarkSuccess()
			}
			return response, err
		}
	}
}
