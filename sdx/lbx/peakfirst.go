package lbx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// PeakFirstFactory create a peak first balancer
type PeakFirstFactory struct{}

func (PeakFirstFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return BalancerFunc(func() (endpoint.Endpoint, error) {
		endpoints, err := endpointer.Endpoints()
		if err != nil {
			return nil, err
		}
		if len(endpoints) <= 0 {
			return nil, lb.ErrNoEndpoints
		}
		return endpoints[0], nil
	})
}
