// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/cqrs/v1"
)

var _ CommandHandler = (*commandHandler)(nil)

type CommandHandler cqrs.CommandHandler[CommandCommand]

type CommandCommand struct {
	v1.UnimplementedCommandCommand
}

func NewCommandHandler() CommandHandler {
	return &commandHandler{}
}

type commandHandler struct {
}

func (h *commandHandler) Handle(ctx context.Context, cmd CommandCommand) error {
	// TODO implement me
	panic("implement me")
}
