package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NamedResponseArgs struct {
}

type NamedResponseRes struct {
}

type NamedResponse cqrs.QueryHandler[*NamedResponseArgs, *NamedResponseRes]

func NewNamedResponse() NamedResponse {
	return &namedResponse{}
}

type namedResponse struct {
}

func (h *namedResponse) Handle(ctx context.Context, args *NamedResponseArgs) (*NamedResponseRes, error) {
	// TODO implement me
	panic("implement me")
}
