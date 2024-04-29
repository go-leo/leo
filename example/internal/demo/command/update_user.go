package command

import (
	"context"
	"fmt"
	"github.com/go-leo/cqrs"
	"github.com/go-leo/cqrs/example/internal/demo/model"
)

type UpdateUserArgs struct {
	User *model.User
}

type UpdateUser cqrs.CommandHandler[*UpdateUserArgs]

func NewUpdateUser() UpdateUser {
	return &updateUser{}
}

type updateUser struct {
}

func (h *updateUser) Handle(ctx context.Context, args *UpdateUserArgs) error {
	fmt.Println(args.User)
	return nil
}