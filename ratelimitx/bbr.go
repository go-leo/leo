package ratelimitx

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	kratosrate "github.com/go-kratos/aegis/ratelimit"
	"github.com/go-kratos/aegis/ratelimit/bbr"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
)

// BBR creates a rate-limiting middleware using a bbrLimiter.
// see: https://github.com/alibaba/Sentinel/wiki/系统自适应限流
// see: https://go-kratos.dev/docs/component/middleware/ratelimit/
// see: https://github.com/go-kratos/aegis/tree/main/ratelimit/bbr
func BBR(limiter *bbr.BBR) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			// extract endpoint name
			endpointName, ok := endpointx.NameExtractor(ctx)
			if !ok {
				return next(ctx, request)
			}
			// limit the request
			doneFunc, err := limiter.Allow()
			if err != nil {
				return nil, statusx.ResourceExhausted(statusx.Message(errMessageFormat, endpointName))
			}
			// continue handle the request
			response, err := next(ctx, request)
			doneFunc(kratosrate.DoneInfo{Err: err})
			return response, err
		}
	}
}
