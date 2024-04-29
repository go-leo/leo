package command

import (
	"context"
	"fmt"
	"github.com/go-leo/cqrs"
)

type DeleteUserArgs struct {
	Name string
}

type DeleteUser cqrs.CommandHandler[*DeleteUserArgs]

func NewDeleteUser() DeleteUser {
	return &deleteUser{}
}

type deleteUser struct {
}

func (h *deleteUser) Handle(ctx context.Context, args *DeleteUserArgs) error {
	fmt.Println("delete user", args)
	return nil
}
