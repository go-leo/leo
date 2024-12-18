package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type EmbedNamedPathOptStringArgs struct {
}

type EmbedNamedPathOptStringRes struct {
}

type EmbedNamedPathOptString cqrs.QueryHandler[*EmbedNamedPathOptStringArgs, *EmbedNamedPathOptStringRes]

func NewEmbedNamedPathOptString() EmbedNamedPathOptString {
	return &embedNamedPathOptString{}
}

type embedNamedPathOptString struct {
}

func (h *embedNamedPathOptString) Handle(ctx context.Context, args *EmbedNamedPathOptStringArgs) (*EmbedNamedPathOptStringRes, error) {
	// TODO implement me
	panic("implement me")
}
