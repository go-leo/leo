package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type FloatPathArgs struct {
}

type FloatPathRes struct {
}

type FloatPath cqrs.QueryHandler[*FloatPathArgs, *FloatPathRes]

func NewFloatPath() FloatPath {
	return &floatPath{}
}

type floatPath struct {
}

func (h *floatPath) Handle(ctx context.Context, args *FloatPathArgs) (*FloatPathRes, error) {
	// TODO implement me
	panic("implement me")
}
