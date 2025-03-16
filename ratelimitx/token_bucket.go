package ratelimitx

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

// TokenBucket accepts a parameter limiter of type *rate.Limiter and returns an endpoint.Middleware. It converts the
// rate limiter into a middleware by calling the waiterMiddleware function, which is used to control the request rate.
// see: https://pkg.go.dev/golang.org/x/time/rate
func TokenBucket(limiter *rate.Limiter) endpoint.Middleware {
	return allowerMiddleware(limiter)
}
