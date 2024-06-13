package httpx

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/sd"
	httptransport "github.com/go-kit/kit/transport/http"
	"io"
)

func Factory(
	scheme string,
	req func(scheme string, instance string) httptransport.CreateRequestFunc,
	dec httptransport.DecodeResponseFunc,
	options ...httptransport.ClientOption,
) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		client := httptransport.NewExplicitClient(req(scheme, instance), dec, options...)
		return client.Endpoint(), io.NopCloser(nil), nil
	}
}
