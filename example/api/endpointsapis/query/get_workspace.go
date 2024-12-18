package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type GetWorkspaceArgs struct {
}

type GetWorkspaceRes struct {
}

type GetWorkspace cqrs.QueryHandler[*GetWorkspaceArgs, *GetWorkspaceRes]

func NewGetWorkspace() GetWorkspace {
	return &getWorkspace{}
}

type getWorkspace struct {
}

func (h *getWorkspace) Handle(ctx context.Context, args *GetWorkspaceArgs) (*GetWorkspaceRes, error) {
	// TODO implement me
	panic("implement me")
}
