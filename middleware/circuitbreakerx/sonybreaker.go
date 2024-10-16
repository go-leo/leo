package circuitbreakerx

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/sony/gobreaker"
)

type GoBreaker struct {
	group lazyloadx.Group
}

func NewGoBreaker(factory func() *gobreaker.CircuitBreaker) *GoBreaker {
	return &GoBreaker{group: lazyloadx.Group{New: func() (any, error) { return factory(), nil }}}
}

func (cb *GoBreaker) Execute(ctx context.Context, request any, endpointName string, next endpoint.Endpoint) (any, error, bool) {
	value, err, _ := cb.group.Load(endpointName)
	if err != nil {
		panic(fmt.Errorf("circuitbreakerx: failed to load %s breaker", endpointName))
	}
	breaker := value.(*gobreaker.CircuitBreaker)
	response, err := breaker.Execute(func() (any, error) {
		return next(ctx, request)
	})
	if errors.Is(err, gobreaker.ErrTooManyRequests) {
		return nil, nil, false
	}
	if errors.Is(err, gobreaker.ErrOpenState) {
		return nil, nil, false
	}
	return response, err, true
}
