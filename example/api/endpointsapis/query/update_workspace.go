package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type UpdateWorkspaceArgs struct {
}

type UpdateWorkspaceRes struct {
}

type UpdateWorkspace cqrs.QueryHandler[*UpdateWorkspaceArgs, *UpdateWorkspaceRes]

func NewUpdateWorkspace() UpdateWorkspace {
	return &updateWorkspace{}
}

type updateWorkspace struct {
}

func (h *updateWorkspace) Handle(ctx context.Context, args *UpdateWorkspaceArgs) (*UpdateWorkspaceRes, error) {
	// TODO implement me
	panic("implement me")
}
