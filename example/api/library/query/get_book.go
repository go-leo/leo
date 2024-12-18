package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetBookArgs struct {
}

type GetBookRes struct {
}

type GetBook cqrs.QueryHandler[*GetBookArgs, *GetBookRes]

func NewGetBook() GetBook {
	return &getBook{}
}

type getBook struct {
}

func (h *getBook) Handle(ctx context.Context, args *GetBookArgs) (*GetBookRes, error) {
	// TODO implement me
	panic("implement me")
}
