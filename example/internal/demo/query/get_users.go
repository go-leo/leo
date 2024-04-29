package query

import (
	"context"
	"github.com/go-leo/cqrs"
	"github.com/go-leo/cqrs/example/internal/demo/model"
	"github.com/go-leo/gox/mathx/randx"
)

type GetUsersArgs struct {
	PageNo   int32
	PageSize int32
}

type GetUsersRes struct {
	List []*model.User
}

type GetUsers cqrs.QueryHandler[*GetUsersArgs, *GetUsersRes]

func NewGetUsers() GetUsers {
	return &getUsers{}
}

type getUsers struct {
}

func (h *getUsers) Handle(ctx context.Context, args *GetUsersArgs) (*GetUsersRes, error) {
	users := make([]*model.User, 0)
	for i := 0; i < int(args.PageSize); i++ {
		users = append(users, &model.User{
			Name:   randx.HexString(12),
			Age:    randx.Int31n(50),
			Salary: float64(randx.Int31n(30000)),
			Token:  randx.NumericString(16),
		})
	}
	return &GetUsersRes{List: users}, nil
}
