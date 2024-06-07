package circuitbreakerx

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/endpointx"
)

// HystrixNewer is the Breaker factory.
type HystrixNewer struct{}

func (HystrixNewer) New() Breaker {
	return &HystrixBreaker{}
}

// HystrixBreaker is the Breaker implementation.
type HystrixBreaker struct{}

func (breaker *HystrixBreaker) Execute(ctx context.Context, request any, next endpoint.Endpoint) (bool, any, error) {
	name, ok := endpointx.ExtractName(ctx)
	if !ok {
		resp, err := next(ctx, request)
		return true, resp, err
	}
	var resp interface{}
	runFunc := func() error {
		var err error
		resp, err = next(ctx, request)
		return err
	}
	if err := hystrix.Do(name, runFunc, nil); err != nil {
		var circuitErr hystrix.CircuitError
		if errors.As(err, &circuitErr) {
			return false, resp, err
		}
		return true, nil, err
	}
	return true, resp, nil
}
