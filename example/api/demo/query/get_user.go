package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetUserArgs struct {
}

type GetUserRes struct {
}

type GetUser cqrs.QueryHandler[*GetUserArgs, *GetUserRes]

func NewGetUser() GetUser {
	return &getUser{}
}

type getUser struct {
}

func (h *getUser) Handle(ctx context.Context, args *GetUserArgs) (*GetUserRes, error) {
	// TODO implement me
	panic("implement me")
}
