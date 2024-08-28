package lbx

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// LeastRequestFactory create a least request balancer
type LeastRequestFactory struct{}

func (f LeastRequestFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return &leastRequestBalancer{
		Ctx:        ctx,
		Endpointer: endpointer,
	}
}

type leastRequestBalancer struct {
	Ctx        context.Context
	Endpointer sd.Endpointer
}

func (b *leastRequestBalancer) Endpoint() (endpoint.Endpoint, error) {
	endpoints, err := b.Endpointer.Endpoints()
	if err != nil {
		return nil, err
	}
	length := uint(len(endpoints))
	if length <= 0 {
		return nil, lb.ErrNoEndpoints
	}
	fmt.Println("leastRequestBalancer")
	return func(ctx context.Context, request any) (response any, err error) {
		fmt.Println("leastRequestBalancer Endpoint")
		return endpoints[0](ctx, request)
	}, nil
}
