package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type ListBooksArgs struct {
}

type ListBooksRes struct {
}

type ListBooks cqrs.QueryHandler[*ListBooksArgs, *ListBooksRes]

func NewListBooks() ListBooks {
	return &listBooks{}
}

type listBooks struct {
}

func (h *listBooks) Handle(ctx context.Context, args *ListBooksArgs) (*ListBooksRes, error) {
	// TODO implement me
	panic("implement me")
}
