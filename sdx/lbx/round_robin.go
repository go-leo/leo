package lbx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"sync/atomic"
)

// RoundRobinFactory create a round robin balancer
type RoundRobinFactory struct {
}

func (f RoundRobinFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRoundRobin(endpointer)
}

// 加权轮询

type WeightedRoundRobinFactory struct {
}

func (f WeightedRoundRobinFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return &weightedRoundRobin{
		s: endpointer,
		c: 0,
	}
}

type weightedRoundRobin struct {
	s sd.Endpointer
	c uint64
}

func (wrr *weightedRoundRobin) Endpoint() (endpoint.Endpoint, error) {
	endpoints, err := wrr.s.Endpoints()
	if err != nil {
		return nil, err
	}
	if len(endpoints) <= 0 {
		return nil, lb.ErrNoEndpoints
	}
	old := atomic.AddUint64(&wrr.c, 1) - 1
	idx := old % uint64(len(endpoints))
	return endpoints[idx], nil
}