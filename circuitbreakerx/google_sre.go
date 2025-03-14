package circuitbreakerx

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kratos/aegis/circuitbreaker"
	"github.com/go-leo/gox/syncx/lazyloadx"
	"github.com/go-leo/leo/v3/endpointx"
	"github.com/go-leo/leo/v3/statusx"
	"google.golang.org/grpc/codes"
)

// GoogleSreBreaker
// see: github.com/go-kratos/aegis/circuitbreaker/sre
// see: https://landing.google.com/sre/sre-book/chapters/handling-overload/
func GoogleSreBreaker(factory func(endpointName string) (circuitbreaker.CircuitBreaker, error)) endpoint.Middleware {
	group := lazyloadx.Group[circuitbreaker.CircuitBreaker]{
		New: factory,
	}
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request any) (any, error) {
			endpointName, ok := endpointx.NameExtractor(ctx)
			if !ok {
				return next(ctx, request)
			}
			cb, err, _ := group.Load(endpointName)
			if err != nil {
				return nil, statusx.Canceled(statusx.Message(errLoadBreaker, endpointName))
			}
			if err := cb.Allow(); errors.Is(err, circuitbreaker.ErrNotAllowed) {
				cb.MarkFailed()
				return nil, ErrCircuitOpen
			}
			response, err := next(ctx, request)
			if err == nil {
				cb.MarkSuccess()
				return response, err
			}
			st, ok := statusx.From(err)
			if !ok {
				cb.MarkSuccess()
				return response, err
			}
			if st.Code() == codes.Unknown ||
				st.Code() == codes.Internal ||
				st.Code() == codes.DataLoss ||
				st.Code() == codes.Unavailable {
				cb.MarkFailed()
				return response, err
			}
			cb.MarkSuccess()
			return response, err
		}
	}
}
