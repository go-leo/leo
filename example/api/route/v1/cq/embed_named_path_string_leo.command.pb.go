// Code generated by protoc-gen-go-leo. DO NOT EDIT.
// If you want edit it, can move this file to another directory.

package cq

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
	v1 "github.com/go-leo/leo/v3/example/api/route/v1"
)

var _ EmbedNamedPathStringHandler = (*embedNamedPathStringHandler)(nil)

type EmbedNamedPathStringHandler cqrs.CommandHandler[EmbedNamedPathStringCommand]

type EmbedNamedPathStringCommand struct {
	v1.UnimplementedEmbedNamedPathStringCommand
}

func (EmbedNamedPathStringCommand) From(ctx context.Context, req *v1.EmbedNamedPathRequest) (v1.EmbedNamedPathStringCommand, context.Context, error) {
	panic("implement me")
	return EmbedNamedPathStringCommand{}, ctx, nil
}

func NewEmbedNamedPathStringHandler() EmbedNamedPathStringHandler {
	return &embedNamedPathStringHandler{}
}

type embedNamedPathStringHandler struct {
}

func (h *embedNamedPathStringHandler) Handle(ctx context.Context, cmd EmbedNamedPathStringCommand) error {
	// TODO implement me
	panic("implement me")
}
