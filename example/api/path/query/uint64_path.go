package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type Uint64PathArgs struct {
}

type Uint64PathRes struct {
}

type Uint64Path cqrs.QueryHandler[*Uint64PathArgs, *Uint64PathRes]

func NewUint64Path() Uint64Path {
	return &uint64Path{}
}

type uint64Path struct {
}

func (h *uint64Path) Handle(ctx context.Context, args *Uint64PathArgs) (*Uint64PathRes, error) {
	// TODO implement me
	panic("implement me")
}
