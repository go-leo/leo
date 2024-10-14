package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/endpointx"
)

// HystrixBreaker is the Breaker implementation.
type HystrixBreaker struct{}

func (breaker *HystrixBreaker) Execute(ctx context.Context, request any, next endpoint.Endpoint) (any, bool, error) {
	name, ok := endpointx.ExtractName(ctx)
	if !ok {
		resp, err := next(ctx, request)
		return resp, true, err
	}
	var resp any
	runFunc := func() error {
		var err error
		resp, err = next(ctx, request)
		return err
	}
	if err := hystrix.Do(name, runFunc, nil); err != nil {
		var circuitErr hystrix.CircuitError
		if errors.As(err, &circuitErr) {
			return resp, false, err
		}
		return nil, true, err
	}
	return resp, true, nil
}
