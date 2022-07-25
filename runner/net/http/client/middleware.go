package client

import (
	"context"
	"net/http"
)

type HTTPInfo struct {
	Path     string
	Client   *http.Client
	Request  *http.Request
	Response *http.Response
}

type Invoker func(ctx context.Context, req any, reply any, info *HTTPInfo) error

type Interceptor func(ctx context.Context, req any, reply any, info *HTTPInfo, invoke Invoker) error

func Chain(middlewares ...Interceptor) Interceptor {
	if len(middlewares) == 0 {
		return nil
	} else if len(middlewares) == 1 {
		return middlewares[0]
	} else {
		return func(ctx context.Context, req any, reply any, info *HTTPInfo, invoke Invoker) error {
			var i int
			var next Invoker
			next = func(ctx context.Context, req any, reply any, info *HTTPInfo) error {
				if i == len(middlewares)-1 {
					return middlewares[i](ctx, req, reply, info, invoke)
				}
				i++
				return middlewares[i-1](ctx, req, reply, info, next)
			}
			return next(ctx, req, reply, info)
		}
	}
}
