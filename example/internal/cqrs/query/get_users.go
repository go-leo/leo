package query

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetUsersArgs struct {
}

type GetUsersRes struct {
}

type GetUsers cqrs.QueryHandler[*GetUsersArgs, *GetUsersRes]

func NewGetUsers() GetUsers {
	return &getUsers{}
}

type getUsers struct {
}

func (h *getUsers) Handle(ctx context.Context, args *GetUsersArgs) (*GetUsersRes, error) {
	// TODO implement me
	panic("implement me")
}
