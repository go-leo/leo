// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/demo/v1"
)

var _ DeleteUserHandler = (*deleteUserHandler)(nil)

type DeleteUserHandler cqrs.CommandHandler[DeleteUserCommand]

type DeleteUserCommand struct {
	v1.UnimplementedDeleteUserCommand
}

func (DeleteUserCommand) From(ctx context.Context, req *v1.DeleteUsersRequest) (v1.DeleteUserCommand, context.Context, error) {
	panic("implement me")
	return DeleteUserCommand{}, ctx, nil
}

func NewDeleteUserHandler() DeleteUserHandler {
	return &deleteUserHandler{}
}

type deleteUserHandler struct {
}

func (h *deleteUserHandler) Handle(ctx context.Context, cmd DeleteUserCommand) error {
	// TODO implement me
	panic("implement me")
}
