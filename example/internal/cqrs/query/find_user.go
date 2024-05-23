package query

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type FindUserArgs struct {
}

type FindUserRes struct {
}

type FindUser cqrs.QueryHandler[*FindUserArgs, *FindUserRes]

func NewFindUser() FindUser {
	return &findUser{}
}

type findUser struct {
}

func (h *findUser) Handle(ctx context.Context, args *FindUserArgs) (*FindUserRes, error) {
	// TODO implement me
	panic("implement me")
}
