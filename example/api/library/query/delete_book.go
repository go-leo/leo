package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DeleteBookArgs struct {
}

type DeleteBookRes struct {
}

type DeleteBook cqrs.QueryHandler[*DeleteBookArgs, *DeleteBookRes]

func NewDeleteBook() DeleteBook {
	return &deleteBook{}
}

type deleteBook struct {
}

func (h *deleteBook) Handle(ctx context.Context, args *DeleteBookArgs) (*DeleteBookRes, error) {
	// TODO implement me
	panic("implement me")
}
