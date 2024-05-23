package query

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
)

type GetUserArgs struct {
	UserId uint64
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
	fmt.Println("get user", args)
	return &GetUserRes{User: &model.User{
		UserId: args.UserId,
		Name:   randx.HexString(8),
		Age:    randx.Int31(),
		Salary: randx.Float64(),
		Token:  randx.NumericString(32),
		Avatar: randx.WordString(16),
	}}, nil
}
