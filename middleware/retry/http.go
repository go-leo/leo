package circuitbreaker

import (
	"context"

	"github.com/go-leo/backoffx"

	"github.com/go-leo/retryx"

	"github.com/go-leo/leo/runner/net/http/client"
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
