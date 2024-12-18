package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type CreateWorkspaceArgs struct {
}

type CreateWorkspaceRes struct {
}

type CreateWorkspace cqrs.QueryHandler[*CreateWorkspaceArgs, *CreateWorkspaceRes]

func NewCreateWorkspace() CreateWorkspace {
	return &createWorkspace{}
}

type createWorkspace struct {
}

func (h *createWorkspace) Handle(ctx context.Context, args *CreateWorkspaceArgs) (*CreateWorkspaceRes, error) {
	// TODO implement me
	panic("implement me")
}
