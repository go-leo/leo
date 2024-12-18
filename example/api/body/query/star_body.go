package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type StarBodyArgs struct {
}

type StarBodyRes struct {
}

type StarBody cqrs.QueryHandler[*StarBodyArgs, *StarBodyRes]

func NewStarBody() StarBody {
	return &starBody{}
}

type starBody struct {
}

func (h *starBody) Handle(ctx context.Context, args *StarBodyArgs) (*StarBodyRes, error) {
	// TODO implement me
	panic("implement me")
}
