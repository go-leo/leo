// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/route/v1"
)

var _ Int32QueryHandler = (*int32QueryHandler)(nil)

type Int32QueryHandler cqrs.CommandHandler[Int32QueryCommand]

type Int32QueryCommand struct {
	v1.UnimplementedInt32QueryCommand
}

func (Int32QueryCommand) From(ctx context.Context, req *v1.Int32QueryRequest) (v1.Int32QueryCommand, context.Context, error) {
	panic("implement me")
	return Int32QueryCommand{}, ctx, nil
}

func NewInt32QueryHandler() Int32QueryHandler {
	return &int32QueryHandler{}
}

type int32QueryHandler struct {
}

func (h *int32QueryHandler) Handle(ctx context.Context, cmd Int32QueryCommand) error {
	// TODO implement me
	panic("implement me")
}
