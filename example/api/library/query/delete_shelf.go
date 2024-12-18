package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DeleteShelfArgs struct {
}

type DeleteShelfRes struct {
}

type DeleteShelf cqrs.QueryHandler[*DeleteShelfArgs, *DeleteShelfRes]

func NewDeleteShelf() DeleteShelf {
	return &deleteShelf{}
}

type deleteShelf struct {
}

func (h *deleteShelf) Handle(ctx context.Context, args *DeleteShelfArgs) (*DeleteShelfRes, error) {
	// TODO implement me
	panic("implement me")
}
