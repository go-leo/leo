package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type MixPathArgs struct {
}

type MixPathRes struct {
}

type MixPath cqrs.QueryHandler[*MixPathArgs, *MixPathRes]

func NewMixPath() MixPath {
	return &mixPath{}
}

type mixPath struct {
}

func (h *mixPath) Handle(ctx context.Context, args *MixPathArgs) (*MixPathRes, error) {
	// TODO implement me
	panic("implement me")
}
