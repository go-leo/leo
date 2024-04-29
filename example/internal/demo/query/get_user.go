package query

import (
	"context"
	"github.com/go-leo/cqrs"
	"github.com/go-leo/cqrs/example/internal/demo/model"
)

type GetUserArgs struct {
	User *model.User
}

type GetUserRes struct {
	User *model.User
}

type GetUser cqrs.QueryHandler[*GetUserArgs, *GetUserRes]

func NewGetUser() GetUser {
	return &getUser{}
}

type getUser struct {
}

func (h *getUser) Handle(ctx context.Context, args *GetUserArgs) (*GetUserRes, error) {
	return &GetUserRes{User: args.User}, nil
}
