package lbx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

type BalancerFactory interface {
	New(ctx context.Context, Endpointer sd.Endpointer) lb.Balancer
}

type RandomFactory struct {
	Endpointer sd.Endpointer
	Seed       int64
}

func (f RandomFactory) New(ctx context.Context, Endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRandom(f.Endpointer, f.Seed)
}

type RoundRobinFactory struct {
	Endpointer sd.Endpointer
}

func (f RoundRobinFactory) New(ctx context.Context, Endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRoundRobin(f.Endpointer)
}
