package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetUserAvatarArgs struct {
}

type GetUserAvatarRes struct {
}

type GetUserAvatar cqrs.QueryHandler[*GetUserAvatarArgs, *GetUserAvatarRes]

func NewGetUserAvatar() GetUserAvatar {
	return &getUserAvatar{}
}

type getUserAvatar struct {
}

func (h *getUserAvatar) Handle(ctx context.Context, args *GetUserAvatarArgs) (*GetUserAvatarRes, error) {
	// TODO implement me
	panic("implement me")
}
