package lbx

import (
	"context"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

type BalancerFactory interface {
	New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer
}

type RandomFactory struct {
	Seed int64
}

func (f RandomFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRandom(endpointer, f.Seed)
}

type RoundRobinFactory struct {
}

func (f RoundRobinFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRoundRobin(endpointer)
}

// p2c 算法
