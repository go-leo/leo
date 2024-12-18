package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type MergeShelvesArgs struct {
}

type MergeShelvesRes struct {
}

type MergeShelves cqrs.QueryHandler[*MergeShelvesArgs, *MergeShelvesRes]

func NewMergeShelves() MergeShelves {
	return &mergeShelves{}
}

type mergeShelves struct {
}

func (h *mergeShelves) Handle(ctx context.Context, args *MergeShelvesArgs) (*MergeShelvesRes, error) {
	// TODO implement me
	panic("implement me")
}
