package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NamedBodyArgs struct {
}

type NamedBodyRes struct {
}

type NamedBody cqrs.QueryHandler[*NamedBodyArgs, *NamedBodyRes]

func NewNamedBody() NamedBody {
	return &namedBody{}
}

type namedBody struct {
}

func (h *namedBody) Handle(ctx context.Context, args *NamedBodyArgs) (*NamedBodyRes, error) {
	// TODO implement me
	panic("implement me")
}
