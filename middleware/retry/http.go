package circuitbreaker

import (
	"context"

	"github.com/hmldd/leo/common/backoffx"
	"github.com/hmldd/leo/common/retryx"
	"github.com/hmldd/leo/runner/net/http/client"
)

func HTTPClientMiddleware(maxAttempts uint, backoffFunc backoffx.BackoffFunc, isRetriable func(error) bool) client.Interceptor {
	return func(ctx context.Context, req any, reply any, info *client.HTTPInfo, invoke client.Invoker) error {
		return retryx.Call(ctx, maxAttempts, backoffFunc, func(_ int) error {
			err := invoke(ctx, req, reply, info)
			if err != nil && isRetriable(err) {
				return err
			}
			return nil
		})
	}
}
