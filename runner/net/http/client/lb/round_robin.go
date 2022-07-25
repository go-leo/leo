package lb

import (
	"context"
	"errors"
	"sync/atomic"
)

type RoundRobin struct {
	number uint64
}

func (r *RoundRobin) Pick(_ context.Context, takers []TargetTaker) (TargetTaker, error) {
	if len(takers) <= 0 {
		return nil, errors.New("takers is empty")
	}
	number := atomic.AddUint64(&r.number, 1)
	index := number % uint64(len(takers))
	return takers[index], nil
}
