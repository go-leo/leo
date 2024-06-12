package httpx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type Client struct {
	client *httptransport.Client
}

func NewClient(req httptransport.CreateRequestFunc, dec httptransport.DecodeResponseFunc, options ...httptransport.ClientOption) *Client {
	return &Client{client: httptransport.NewExplicitClient(req, dec, options...)}
}

func (c *Client) Endpoint(ctx context.Context) endpoint.Endpoint {
	return c.client.Endpoint()
}
