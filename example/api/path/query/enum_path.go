package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type EnumPathArgs struct {
}

type EnumPathRes struct {
}

type EnumPath cqrs.QueryHandler[*EnumPathArgs, *EnumPathRes]

func NewEnumPath() EnumPath {
	return &enumPath{}
}

type enumPath struct {
}

func (h *enumPath) Handle(ctx context.Context, args *EnumPathArgs) (*EnumPathRes, error) {
	// TODO implement me
	panic("implement me")
}
