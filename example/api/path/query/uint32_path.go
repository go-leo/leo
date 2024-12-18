package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type Uint32PathArgs struct {
}

type Uint32PathRes struct {
}

type Uint32Path cqrs.QueryHandler[*Uint32PathArgs, *Uint32PathRes]

func NewUint32Path() Uint32Path {
	return &uint32Path{}
}

type uint32Path struct {
}

func (h *uint32Path) Handle(ctx context.Context, args *Uint32PathArgs) (*Uint32PathRes, error) {
	// TODO implement me
	panic("implement me")
}
