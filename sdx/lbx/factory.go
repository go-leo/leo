package lbx

import (
	"context"

	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// BalancerFactory create a balancer
type BalancerFactory interface {
	New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer
}

// p2c 算法
