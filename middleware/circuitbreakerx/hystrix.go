package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-kit/kit/endpoint"
)

// HystrixFactory is the Breaker factory.
type HystrixFactory struct{}

func (HystrixFactory) Create() Breaker { return &HystrixBreaker{} }

// HystrixBreaker is the Breaker implementation.
type HystrixBreaker struct{}

func (breaker *HystrixBreaker) Execute(ctx context.Context, request any, endpointName string, next endpoint.Endpoint) (any, error, bool) {
	var resp any
	runFunc := func() error {
		var err error
		resp, err = next(ctx, request)
		return err
	}
	if err := hystrix.Do(endpointName, runFunc, nil); err != nil {
		var circuitErr hystrix.CircuitError
		if errors.As(err, &circuitErr) {
			return resp, nil, false
		}
		return nil, err, true
	}
	return resp, nil, true
}
