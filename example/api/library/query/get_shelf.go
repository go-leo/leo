package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetShelfArgs struct {
}

type GetShelfRes struct {
}

type GetShelf cqrs.QueryHandler[*GetShelfArgs, *GetShelfRes]

func NewGetShelf() GetShelf {
	return &getShelf{}
}

type getShelf struct {
}

func (h *getShelf) Handle(ctx context.Context, args *GetShelfArgs) (*GetShelfRes, error) {
	// TODO implement me
	panic("implement me")
}
