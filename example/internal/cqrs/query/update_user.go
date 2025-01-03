package query

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type UpdateUserArgs struct {
}

type UpdateUserRes struct {
}

type UpdateUser cqrs.QueryHandler[*UpdateUserArgs, *UpdateUserRes]

func NewUpdateUser() UpdateUser {
	return &updateUser{}
}

type updateUser struct {
}

func (h *updateUser) Handle(ctx context.Context, args *UpdateUserArgs) (*UpdateUserRes, error) {
	// TODO implement me
	panic("implement me")
}
