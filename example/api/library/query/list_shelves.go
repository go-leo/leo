package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type ListShelvesArgs struct {
}

type ListShelvesRes struct {
}

type ListShelves cqrs.QueryHandler[*ListShelvesArgs, *ListShelvesRes]

func NewListShelves() ListShelves {
	return &listShelves{}
}

type listShelves struct {
}

func (h *listShelves) Handle(ctx context.Context, args *ListShelvesArgs) (*ListShelvesRes, error) {
	// TODO implement me
	panic("implement me")
}
