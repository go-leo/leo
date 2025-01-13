package command

import (
	"context"
	"github.com/go-leo/leo/v3/cqrs"
)

type CreateUserArgs struct {
}

type CreateUser cqrs.CommandHandler[*CreateUserArgs]

func NewCreateUser() CreateUser {
	return &createUser{}
}

type createUser struct {
}

func (h *createUser) Handle(ctx context.Context, args *CreateUserArgs) error {
	// TODO implement me
	panic("implement me")
}
