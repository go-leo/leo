package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NamedPathWrapStringArgs struct {
}

type NamedPathWrapStringRes struct {
}

type NamedPathWrapString cqrs.QueryHandler[*NamedPathWrapStringArgs, *NamedPathWrapStringRes]

func NewNamedPathWrapString() NamedPathWrapString {
	return &namedPathWrapString{}
}

type namedPathWrapString struct {
}

func (h *namedPathWrapString) Handle(ctx context.Context, args *NamedPathWrapStringArgs) (*NamedPathWrapStringRes, error) {
	// TODO implement me
	panic("implement me")
}
