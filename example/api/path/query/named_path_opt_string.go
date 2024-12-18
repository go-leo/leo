package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NamedPathOptStringArgs struct {
}

type NamedPathOptStringRes struct {
}

type NamedPathOptString cqrs.QueryHandler[*NamedPathOptStringArgs, *NamedPathOptStringRes]

func NewNamedPathOptString() NamedPathOptString {
	return &namedPathOptString{}
}

type namedPathOptString struct {
}

func (h *namedPathOptString) Handle(ctx context.Context, args *NamedPathOptStringArgs) (*NamedPathOptStringRes, error) {
	// TODO implement me
	panic("implement me")
}
