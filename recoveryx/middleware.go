package recoveryx

import (
	"context"
	"runtime"
	"strings"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/convx"
	"github.com/go-leo/status"
	"google.golang.org/grpc/codes"
)

type options struct {
	handler HandlerFunc
}
type Option func(*options)

func (o *options) apply(opts ...Option) *options {
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type HandlerFunc func(ctx context.Context, p any) (err error)

// RecoveryHandler customizes the function for recovering from a panic.
func RecoveryHandler(f HandlerFunc) Option {
	return func(o *options) {
		o.handler = f
	}
}

func Middleware(opts ...Option) endpoint.Middleware {
	opt := new(options).apply(opts...)
	return func(e endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if p := recover(); p != nil {
					if opt.handler != nil {
						err = opt.handler(ctx, p)
						return
					}
					i := 64 << 10
					stack := make([]byte, i)
					stack = stack[:runtime.Stack(stack, false)]
					err = status.New(codes.Internal,
						status.Message("panic triggered: %v", p),
						status.DebugInfo(strings.Split(convx.BytesToString(stack), "\n"), ""),
						status.Identifier("github.com/go-leo/leo/v3/recoveryx.ErrPanicked"),
					)
				}
			}()
			return e(ctx, request)
		}
	}
}
