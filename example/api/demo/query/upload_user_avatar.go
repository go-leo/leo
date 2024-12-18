package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type UploadUserAvatarArgs struct {
}

type UploadUserAvatarRes struct {
}

type UploadUserAvatar cqrs.QueryHandler[*UploadUserAvatarArgs, *UploadUserAvatarRes]

func NewUploadUserAvatar() UploadUserAvatar {
	return &uploadUserAvatar{}
}

type uploadUserAvatar struct {
}

func (h *uploadUserAvatar) Handle(ctx context.Context, args *UploadUserAvatarArgs) (*UploadUserAvatarRes, error) {
	// TODO implement me
	panic("implement me")
}
