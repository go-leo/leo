package lbx

import (
	"context"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
)

// RandomFactory create a random balancer
type RandomFactory struct {
	Seed int64
}

func (f RandomFactory) New(ctx context.Context, endpointer sd.Endpointer) lb.Balancer {
	return lb.NewRandom(endpointer, f.Seed)
}

// 加权随机
