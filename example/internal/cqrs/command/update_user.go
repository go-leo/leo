package command

import (
	"context"
	"github.com/go-leo/leo/v3/cqrs"
)

type UpdateUserArgs struct {
}

type UpdateUser cqrs.CommandHandler[*UpdateUserArgs]

func NewUpdateUser() UpdateUser {
	return &updateUser{}
}

type updateUser struct {
}

func (h *updateUser) Handle(ctx context.Context, args *UpdateUserArgs) error {
	// TODO implement me
	panic("implement me")
}
