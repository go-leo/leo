package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	metadatax "github.com/go-leo/leo/v3/metadatax"
)

type UpdateUserArgs struct {
}

type UpdateUser cqrs.CommandHandler[*UpdateUserArgs]

func NewUpdateUser() UpdateUser {
	return &updateUser{}
}

type updateUser struct {
}

func (h *updateUser) Handle(ctx context.Context, args *UpdateUserArgs) (metadatax.Metadata, error) {
	// TODO implement me
	panic("implement me")
}
