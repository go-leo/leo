package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NamedPathStringArgs struct {
}

type NamedPathStringRes struct {
}

type NamedPathString cqrs.QueryHandler[*NamedPathStringArgs, *NamedPathStringRes]

func NewNamedPathString() NamedPathString {
	return &namedPathString{}
}

type namedPathString struct {
}

func (h *namedPathString) Handle(ctx context.Context, args *NamedPathStringArgs) (*NamedPathStringRes, error) {
	// TODO implement me
	panic("implement me")
}
