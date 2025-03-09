package ratelimitx

import (
	"github.com/RussellLuo/slidingwindow"
	"github.com/go-kit/kit/endpoint"
)

// SlideWindow accepts a sliding window rate limiter as a parameter and returns a middleware function
// allowerMiddleware(limiter). This middleware is typically used for service endpoints to control the frequency of
// requests using a sliding window algorithm, ensuring the system does not become overloaded due to too many requests.
// see: https://github.com/RussellLuo/slidingwindow
func SlideWindow(limiter *slidingwindow.Limiter) endpoint.Middleware {
	return allowerMiddleware(limiter)
}
