package command

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/cqrs"
	"github.com/go-leo/leo/v3/metadatax"
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

func (h *uploadUserAvatar) Handle(ctx context.Context, args *UploadUserAvatarArgs) (metadatax.Metadata, error) {
	fmt.Println("upload user avatar", args)
	return metadatax.New(), nil
}
