package lbx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// BalancerFactory create a balancer
type BalancerFactory interface {
	New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer
}

type BalancerFunc func() (endpoint.Endpoint, error)

func (f BalancerFunc) Endpoint() (endpoint.Endpoint, error) {
	return f()
}

type BalancerFactoryFunc func(ctx context.Context, endpointer sd.Endpointer) lb.Balancer

func (f BalancerFactoryFunc) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return f(ctx, endpointer)
}

func NewBalancer(ctx context.Context, factory BalancerFactory, newEndpointer func(ctx context.Context, color string) (sd.Endpointer, error)) func(key string) (lb.Balancer, error) {
	return func(color string) (lb.Balancer, error) {
		endpointer, err := newEndpointer(ctx, color)
		if err != nil {
			return nil, err
		}
		balancer := factory.New(ctx, endpointer)
		wrappedErr := BalancerFunc(func() (endpoint.Endpoint, error) {
			ep, err := balancer.Endpoint()
			if err != nil {
				return nil, err
			}
			return ep, nil
		})
		return wrappedErr, nil
	}
}
