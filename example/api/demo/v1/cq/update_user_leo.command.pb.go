// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/demo/v1"
)

var _ UpdateUserHandler = (*updateUserHandler)(nil)

type UpdateUserHandler cqrs.CommandHandler[UpdateUserCommand]

type UpdateUserCommand struct {
	v1.UnimplementedUpdateUserCommand
}

func NewUpdateUserHandler() UpdateUserHandler {
	return &updateUserHandler{}
}

type updateUserHandler struct {
}

func (h *updateUserHandler) Handle(ctx context.Context, cmd UpdateUserCommand) error {
	// TODO implement me
	panic("implement me")
}
