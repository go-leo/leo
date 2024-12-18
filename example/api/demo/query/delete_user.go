package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DeleteUserArgs struct {
}

type DeleteUserRes struct {
}

type DeleteUser cqrs.QueryHandler[*DeleteUserArgs, *DeleteUserRes]

func NewDeleteUser() DeleteUser {
	return &deleteUser{}
}

type deleteUser struct {
}

func (h *deleteUser) Handle(ctx context.Context, args *DeleteUserArgs) (*DeleteUserRes, error) {
	// TODO implement me
	panic("implement me")
}
