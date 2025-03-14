package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/endpointx"
)

// Hystrix circuit breaker
// see: https://godoc.org/github.com/afex/hystrix-go/hystrix
// see: https://github.com/Netflix/Hystrix
func Hystrix() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			endpointName, ok := endpointx.NameExtractor(ctx)
			if !ok {
				return next(ctx, request)
			}
			response, err := circuitbreaker.Hystrix(endpointName)(next)(ctx, request)
			var circuitErr hystrix.CircuitError
			if errors.As(err, &circuitErr) {
				return nil, ErrCircuitOpen
			}
			return response, err
		}
	}
}
