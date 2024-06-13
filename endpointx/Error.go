package endpointx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func Error(err error) endpoint.Endpoint {
	return func(context.Context, any) (any, error) {
		return nil, err
	}
}
