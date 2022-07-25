package lb

import (
	"context"
	"errors"
)

type PickFirst struct{}

func (p *PickFirst) Pick(_ context.Context, takers []TargetTaker) (TargetTaker, error) {
	if len(takers) <= 0 {
		return nil, errors.New("takers is empty")
	}
	return takers[0], nil
}
