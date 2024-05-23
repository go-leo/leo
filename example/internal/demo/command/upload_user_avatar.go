package command

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
)

type UploadUserAvatarArgs struct {
	UserId uint64
	Avatar []byte
}

type UploadUserAvatar cqrs.CommandHandler[*UploadUserAvatarArgs]

func NewUploadUserAvatar() UploadUserAvatar {
	return &uploadUserAvatar{}
}

type uploadUserAvatar struct {
}

func (h *uploadUserAvatar) Handle(ctx context.Context, args *UploadUserAvatarArgs) (cqrs.Metadata, error) {
	fmt.Println("upload user avatar", args)
	return cqrs.NewMetadata(), nil
}
