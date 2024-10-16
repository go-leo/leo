package circuitbreakerx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
)

// Breaker execute the endpoint.
// if the circuit is open, it will return false, nil, nil
// if the circuit is closed, it will execute the endpoint and return true, response, err.
type Breaker interface {
	Execute(ctx context.Context, request any, endpointName string, next endpoint.Endpoint) (any, error, bool)
}

// Factory create the Breaker.
type Factory interface {
	Create() Breaker
}

func Middleware(breaker Breaker) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			transportName, ok := transportx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}
			switch transportName {
			case grpcx.GrpcServer, httpx.HttpServer:
				return next(ctx, request)
			}

			endpointName, ok := endpointx.ExtractName(ctx)
			if !ok {
				return next(ctx, request)
			}

			resp, err, ok := breaker.Execute(ctx, request, endpointName, next)
			if !ok {
				return nil, statusx.ErrUnavailable.With(statusx.Message("circuit breaker is open"))
			}
			return resp, err
		}
	}
}
