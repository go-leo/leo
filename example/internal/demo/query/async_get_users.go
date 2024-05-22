package query

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/example/internal/demo/model"
	"time"
)

type AsyncGetUsersArgs struct {
	PageNo   int32
	PageSize int32
}

type AsyncGetUsersRes struct {
	List []*model.User
}

type AsyncGetUsers cqrs.QueryHandler[*AsyncGetUsersArgs, *AsyncGetUsersRes]

func NewAsyncGetUsers() AsyncGetUsers {
	return &asyncGetUsers{}
}

type asyncGetUsers struct {
}

func (h *asyncGetUsers) Handle(ctx context.Context, args *AsyncGetUsersArgs) (*AsyncGetUsersRes, error) {
	time.Sleep(10 * time.Second)
	users := make([]*model.User, 0)
	for i := 0; i < int(args.PageSize); i++ {
		users = append(users, &model.User{
			Name:   randx.HexString(12),
			Age:    randx.Int31n(50),
			Salary: float64(randx.Int31n(30000)),
			Token:  randx.NumericString(16),
		})
	}
	return &AsyncGetUsersRes{List: users}, nil
}
