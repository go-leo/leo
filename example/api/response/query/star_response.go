package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type StarResponseArgs struct {
}

type StarResponseRes struct {
}

type StarResponse cqrs.QueryHandler[*StarResponseArgs, *StarResponseRes]

func NewStarResponse() StarResponse {
	return &starResponse{}
}

type starResponse struct {
}

func (h *starResponse) Handle(ctx context.Context, args *StarResponseArgs) (*StarResponseRes, error) {
	// TODO implement me
	panic("implement me")
}
