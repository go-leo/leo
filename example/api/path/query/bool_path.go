package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type BoolPathArgs struct {
}

type BoolPathRes struct {
}

type BoolPath cqrs.QueryHandler[*BoolPathArgs, *BoolPathRes]

func NewBoolPath() BoolPath {
	return &boolPath{}
}

type boolPath struct {
}

func (h *boolPath) Handle(ctx context.Context, args *BoolPathArgs) (*BoolPathRes, error) {
	// TODO implement me
	panic("implement me")
}
