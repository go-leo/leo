package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type CreateShelfArgs struct {
}

type CreateShelfRes struct {
}

type CreateShelf cqrs.QueryHandler[*CreateShelfArgs, *CreateShelfRes]

func NewCreateShelf() CreateShelf {
	return &createShelf{}
}

type createShelf struct {
}

func (h *createShelf) Handle(ctx context.Context, args *CreateShelfArgs) (*CreateShelfRes, error) {
	// TODO implement me
	panic("implement me")
}
