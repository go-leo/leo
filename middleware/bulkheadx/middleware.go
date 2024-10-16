package bulkheadx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/syncx/gopher"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
)

func Middleware(pool gopher.Gopher) endpoint.Middleware {
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

			resC := make(chan any, 1)
			errC := make(chan error, 1)
			err := pool.Go(func() {
				resp, err := next(ctx, request)
				if err != nil {
					errC <- err
					return
				}
				resC <- resp
			})
			if err != nil {
				return next, statusx.ErrResourceExhausted.With(statusx.Wrap(err))
			}

			select {
			case resp := <-resC:
				return resp, nil
			case err := <-errC:
				return nil, err
			}
		}
	}
}
