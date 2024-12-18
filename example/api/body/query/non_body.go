package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type NonBodyArgs struct {
}

type NonBodyRes struct {
}

type NonBody cqrs.QueryHandler[*NonBodyArgs, *NonBodyRes]

func NewNonBody() NonBody {
	return &nonBody{}
}

type nonBody struct {
}

func (h *nonBody) Handle(ctx context.Context, args *NonBodyArgs) (*NonBodyRes, error) {
	// TODO implement me
	panic("implement me")
}
