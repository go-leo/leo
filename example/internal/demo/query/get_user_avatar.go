package query

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
)

type GetUserAvatarArgs struct {
	UserId uint64
}

type GetUserAvatarRes struct {
	Data []byte
}

type GetUserAvatar cqrs.QueryHandler[*GetUserAvatarArgs, *GetUserAvatarRes]

func NewGetUserAvatar() GetUserAvatar {
	return &getUserAvatar{}
}

type getUserAvatar struct {
}

func (h *getUserAvatar) Handle(ctx context.Context, args *GetUserAvatarArgs) (*GetUserAvatarRes, error) {
	fmt.Println("get user avatar", args)
	b := make([]byte, 128)
	_, _ = rand.Read(b)
	return &GetUserAvatarRes{Data: b}, nil
}
