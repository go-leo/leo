package command

import (
	"context"
	"fmt"
	"github.com/go-leo/cqrs"
	"github.com/go-leo/cqrs/example/internal/demo/model"
)

type CreateUserArgs struct {
	User *model.User
}

type CreateUser cqrs.CommandHandler[*CreateUserArgs]

func NewCreateUser() CreateUser {
	return &createUser{}
}

type createUser struct {
}

func (h *createUser) Handle(ctx context.Context, args *CreateUserArgs) error {
	fmt.Println("create user", args)
	return nil
}
