package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	metadatax "github.com/go-leo/leo/v3/metadatax"
)

type DeleteUserArgs struct {
}

type DeleteUser cqrs.CommandHandler[*DeleteUserArgs]

func NewDeleteUser() DeleteUser {
	return &deleteUser{}
}

type deleteUser struct {
}

func (h *deleteUser) Handle(ctx context.Context, args *DeleteUserArgs) (metadatax.Metadata, error) {
	// TODO implement me
	panic("implement me")
}
