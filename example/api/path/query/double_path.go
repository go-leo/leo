package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DoublePathArgs struct {
}

type DoublePathRes struct {
}

type DoublePath cqrs.QueryHandler[*DoublePathArgs, *DoublePathRes]

func NewDoublePath() DoublePath {
	return &doublePath{}
}

type doublePath struct {
}

func (h *doublePath) Handle(ctx context.Context, args *DoublePathArgs) (*DoublePathRes, error) {
	// TODO implement me
	panic("implement me")
}
