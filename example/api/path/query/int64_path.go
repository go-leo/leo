package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type Int64PathArgs struct {
}

type Int64PathRes struct {
}

type Int64Path cqrs.QueryHandler[*Int64PathArgs, *Int64PathRes]

func NewInt64Path() Int64Path {
	return &int64Path{}
}

type int64Path struct {
}

func (h *int64Path) Handle(ctx context.Context, args *Int64PathArgs) (*Int64PathRes, error) {
	// TODO implement me
	panic("implement me")
}
