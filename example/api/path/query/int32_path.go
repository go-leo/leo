package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type Int32PathArgs struct {
}

type Int32PathRes struct {
}

type Int32Path cqrs.QueryHandler[*Int32PathArgs, *Int32PathRes]

func NewInt32Path() Int32Path {
	return &int32Path{}
}

type int32Path struct {
}

func (h *int32Path) Handle(ctx context.Context, args *Int32PathArgs) (*Int32PathRes, error) {
	// TODO implement me
	panic("implement me")
}
