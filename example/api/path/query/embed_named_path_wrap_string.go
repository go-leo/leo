package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type EmbedNamedPathWrapStringArgs struct {
}

type EmbedNamedPathWrapStringRes struct {
}

type EmbedNamedPathWrapString cqrs.QueryHandler[*EmbedNamedPathWrapStringArgs, *EmbedNamedPathWrapStringRes]

func NewEmbedNamedPathWrapString() EmbedNamedPathWrapString {
	return &embedNamedPathWrapString{}
}

type embedNamedPathWrapString struct {
}

func (h *embedNamedPathWrapString) Handle(ctx context.Context, args *EmbedNamedPathWrapStringArgs) (*EmbedNamedPathWrapStringRes, error) {
	// TODO implement me
	panic("implement me")
}
