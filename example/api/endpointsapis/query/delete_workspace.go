package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type DeleteWorkspaceArgs struct {
}

type DeleteWorkspaceRes struct {
}

type DeleteWorkspace cqrs.QueryHandler[*DeleteWorkspaceArgs, *DeleteWorkspaceRes]

func NewDeleteWorkspace() DeleteWorkspace {
	return &deleteWorkspace{}
}

type deleteWorkspace struct {
}

func (h *deleteWorkspace) Handle(ctx context.Context, args *DeleteWorkspaceArgs) (*DeleteWorkspaceRes, error) {
	// TODO implement me
	panic("implement me")
}
