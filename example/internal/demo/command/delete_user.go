package command

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/metadatax"
)

type DeleteUserArgs struct {
	UserId uint64
}

type DeleteUser cqrs.CommandHandler[*DeleteUserArgs]

func NewDeleteUser() DeleteUser {
	return &deleteUser{}
}

type deleteUser struct {
}

func (h *deleteUser) Handle(ctx context.Context, args *DeleteUserArgs) (metadatax.Metadata, error) {
	fmt.Println("delete user", args)
	return metadatax.New(), nil
}
