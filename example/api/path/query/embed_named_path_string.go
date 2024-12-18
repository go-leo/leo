package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type EmbedNamedPathStringArgs struct {
}

type EmbedNamedPathStringRes struct {
}

type EmbedNamedPathString cqrs.QueryHandler[*EmbedNamedPathStringArgs, *EmbedNamedPathStringRes]

func NewEmbedNamedPathString() EmbedNamedPathString {
	return &embedNamedPathString{}
}

type embedNamedPathString struct {
}

func (h *embedNamedPathString) Handle(ctx context.Context, args *EmbedNamedPathStringArgs) (*EmbedNamedPathStringRes, error) {
	// TODO implement me
	panic("implement me")
}
