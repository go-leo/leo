package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type CreateBookArgs struct {
}

type CreateBookRes struct {
}

type CreateBook cqrs.QueryHandler[*CreateBookArgs, *CreateBookRes]

func NewCreateBook() CreateBook {
	return &createBook{}
}

type createBook struct {
}

func (h *createBook) Handle(ctx context.Context, args *CreateBookArgs) (*CreateBookRes, error) {
	// TODO implement me
	panic("implement me")
}
