package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type ListWorkspacesArgs struct {
}

type ListWorkspacesRes struct {
}

type ListWorkspaces cqrs.QueryHandler[*ListWorkspacesArgs, *ListWorkspacesRes]

func NewListWorkspaces() ListWorkspaces {
	return &listWorkspaces{}
}

type listWorkspaces struct {
}

func (h *listWorkspaces) Handle(ctx context.Context, args *ListWorkspacesArgs) (*ListWorkspacesRes, error) {
	// TODO implement me
	panic("implement me")
}
