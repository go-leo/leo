package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	metadatax "github.com/go-leo/leo/v3/metadatax"
)

type CreateUserArgs struct {
}

type CreateUser cqrs.CommandHandler[*CreateUserArgs]

func NewCreateUser() CreateUser {
	return &createUser{}
}

type createUser struct {
}

func (h *createUser) Handle(ctx context.Context, args *CreateUserArgs) (metadatax.Metadata, error) {
	// TODO implement me
	panic("implement me")
}
