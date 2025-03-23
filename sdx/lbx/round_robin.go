package lbx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// RoundRobinFactory create a round robin balancer
type RoundRobinFactory struct{}

func (RoundRobinFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRoundRobin(endpointer)
}

// 加权轮询
// 平滑的加权轮询
