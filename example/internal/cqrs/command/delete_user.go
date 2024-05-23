package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DeleteUserArgs struct {
}

type DeleteUser cqrs.CommandHandler[*DeleteUserArgs]

func NewDeleteUser() DeleteUser {
	return &deleteUser{}
}

type deleteUser struct {
}

func (h *deleteUser) Handle(ctx context.Context, args *DeleteUserArgs) (cqrs.Metadata, error) {
	// TODO implement me
	panic("implement me")
}
