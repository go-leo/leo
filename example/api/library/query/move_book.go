package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type MoveBookArgs struct {
}

type MoveBookRes struct {
}

type MoveBook cqrs.QueryHandler[*MoveBookArgs, *MoveBookRes]

func NewMoveBook() MoveBook {
	return &moveBook{}
}

type moveBook struct {
}

func (h *moveBook) Handle(ctx context.Context, args *MoveBookArgs) (*MoveBookRes, error) {
	// TODO implement me
	panic("implement me")
}
