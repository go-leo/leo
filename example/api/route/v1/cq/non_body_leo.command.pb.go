// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/route/v1"
)

var _ NonBodyHandler = (*nonBodyHandler)(nil)

type NonBodyHandler cqrs.CommandHandler[NonBodyCommand]

type NonBodyCommand struct {
	v1.UnimplementedNonBodyCommand
}

func NewNonBodyHandler() NonBodyHandler {
	return &nonBodyHandler{}
}

type nonBodyHandler struct {
}

func (h *nonBodyHandler) Handle(ctx context.Context, cmd NonBodyCommand) error {
	// TODO implement me
	panic("implement me")
}
