package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type UpdateBookArgs struct {
}

type UpdateBookRes struct {
}

type UpdateBook cqrs.QueryHandler[*UpdateBookArgs, *UpdateBookRes]

func NewUpdateBook() UpdateBook {
	return &updateBook{}
}

type updateBook struct {
}

func (h *updateBook) Handle(ctx context.Context, args *UpdateBookArgs) (*UpdateBookRes, error) {
	// TODO implement me
	panic("implement me")
}
