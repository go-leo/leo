package command

import (
	context "context"
	cqrs "github.com/go-leo/leo/v3/cqrs"
)

type OmittedResponseArgs struct {
}

type OmittedResponseRes struct {
}

type OmittedResponse cqrs.QueryHandler[*OmittedResponseArgs, *OmittedResponseRes]

func NewOmittedResponse() OmittedResponse {
	return &omittedResponse{}
}

type omittedResponse struct {
}

func (h *omittedResponse) Handle(ctx context.Context, args *OmittedResponseArgs) (*OmittedResponseRes, error) {
	// TODO implement me
	panic("implement me")
}
