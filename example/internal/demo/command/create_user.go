package command

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
	"github.com/go-leo/leo/v3/metadatax"
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
	metadata := metadatax.New()
	metadata.Set("id", convx.ToString(randx.Uint64()))
	return nil
}
