// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/route/v1"
)

var _ FloatQueryHandler = (*floatQueryHandler)(nil)

type FloatQueryHandler cqrs.CommandHandler[FloatQueryCommand]

type FloatQueryCommand struct {
	v1.UnimplementedFloatQueryCommand
}

func (FloatQueryCommand) From(ctx context.Context, req *v1.FloatQueryRequest) (v1.FloatQueryCommand, context.Context, error) {
	panic("implement me")
	return FloatQueryCommand{}, ctx, nil
}

func NewFloatQueryHandler() FloatQueryHandler {
	return &floatQueryHandler{}
}

type floatQueryHandler struct {
}

func (h *floatQueryHandler) Handle(ctx context.Context, cmd FloatQueryCommand) error {
	// TODO implement me
	panic("implement me")
}
