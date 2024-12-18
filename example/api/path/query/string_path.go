package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type StringPathArgs struct {
}

type StringPathRes struct {
}

type StringPath cqrs.QueryHandler[*StringPathArgs, *StringPathRes]

func NewStringPath() StringPath {
	return &stringPath{}
}

type stringPath struct {
}

func (h *stringPath) Handle(ctx context.Context, args *StringPathArgs) (*StringPathRes, error) {
	// TODO implement me
	panic("implement me")
}
