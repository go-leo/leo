package endpointx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Handler[Req, Resp any] interface {
	Handle(ctx context.Context, request Req) (Resp, error)
}

func HandleEndpoint[Req, Resp any](e Handler[Req, Resp]) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return e.Handle(ctx, request.(Req))
	}
}
