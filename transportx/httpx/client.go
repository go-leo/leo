package httpx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/leo/v3/sdx"
	"io"
)

func ClientFactory(
	req func(scheme string, instance string) httptransport.CreateRequestFunc,
	dec httptransport.DecodeResponseFunc,
	options ...httptransport.ClientOption,
) sdx.Factory {
	return func(ctx context.Context, args any) sd.Factory {
		if args == nil {
			args = "http"
		}
		return func(instance string) (endpoint.Endpoint, io.Closer, error) {
			scheme, ok := args.(string)
			if !ok {
				return nil, nil, errors.New("invalid http factory args")
			}
			client := httptransport.NewExplicitClient(req(scheme, instance), dec, options...)
			return client.Endpoint(), nil, nil
		}
	}

}
