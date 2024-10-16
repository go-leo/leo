package circuitbreakerx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/statusx"
)

// Gutter Breaker

// GoogleSreBreaker
// see: https://landing.google.com/sre/sre-book/chapters/handling-overload/
type GoogleSreBreaker struct {
	group lazyloadx.Group
}

func NewGoogleSreBreaker(factory func() circuitbreaker.CircuitBreaker) *GoogleSreBreaker {
	return &GoogleSreBreaker{group: lazyloadx.Group{New: func() (any, error) { return factory(), nil }}}
}

func (cb *GoogleSreBreaker) Execute(ctx context.Context, request any, endpointName string, next endpoint.Endpoint) (any, error, bool) {
	value, err, _ := cb.group.Load(endpointName)
	if err != nil {
		panic(fmt.Errorf("circuitbreakerx: failed to load %s breaker", endpointName))
	}
	breaker := value.(circuitbreaker.CircuitBreaker)
	if err := breaker.Allow(); err != nil {
		// rejected
		breaker.MarkFailed()
		return nil, nil, false
	}
	// allowed
	reply, err := next(ctx, request)
	if err != nil && (statusx.ErrInternal.Is(err) || statusx.ErrUnavailable.Equals(err) || statusx.ErrDeadlineExceeded.Equals(err)) {
		breaker.MarkFailed()
	} else {
		breaker.MarkSuccess()
	}
	return reply, err, true
}
