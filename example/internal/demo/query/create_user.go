package query

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type CreateUserArgs struct {
}

type CreateUserRes struct {
	UserId uint64
}

type CreateUser cqrs.QueryHandler[*CreateUserArgs, *CreateUserRes]

func NewCreateUser() CreateUser {
	return &createUser{}
}

type createUser struct {
}

func (h *createUser) Handle(ctx context.Context, args *CreateUserArgs) (*CreateUserRes, error) {
	// TODO implement me
	panic("implement me")
}
