// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/library/v1"
)

var _ DeleteShelfHandler = (*deleteShelfHandler)(nil)

type DeleteShelfHandler cqrs.CommandHandler[DeleteShelfCommand]

type DeleteShelfCommand struct {
	v1.UnimplementedDeleteShelfCommand
}

func (DeleteShelfCommand) From(ctx context.Context, req *v1.DeleteShelfRequest) (v1.DeleteShelfCommand, context.Context, error) {
	panic("implement me")
	return DeleteShelfCommand{}, ctx, nil
}

func NewDeleteShelfHandler() DeleteShelfHandler {
	return &deleteShelfHandler{}
}

type deleteShelfHandler struct {
}

func (h *deleteShelfHandler) Handle(ctx context.Context, cmd DeleteShelfCommand) error {
	// TODO implement me
	panic("implement me")
}
