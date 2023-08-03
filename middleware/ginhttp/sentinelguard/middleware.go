package sentinelguard

import (
	"net/http"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/gin-gonic/gin"
)

// Middleware returns new gin.HandlerFunc
// Default resource name is {method}:{path}, such as "GET:/api/users/:id"
// Default block fallback is returning 429 code
// Define your own behavior by setting options
func Middleware(opts ...Option) gin.HandlerFunc {
	options := evaluateOptions(opts)
	return func(c *gin.Context) {
		resource := c.Request.Method + ":" + c.FullPath()

		if options.resourceExtract != nil {
			resource = options.resourceExtract(c)
		}

		entry, err := sentinel.Entry(
			resource,
			sentinel.WithResourceType(base.ResTypeWeb),
			sentinel.WithTrafficType(base.Inbound),
		)
		if err != nil {
			if options.blockFallback != nil {
				options.blockFallback(c)
			} else {
				c.AbortWithStatus(http.StatusTooManyRequests)
			}
			return
		}

		defer entry.Exit()
		c.Next()
	}
}

type (
	Option  func(*options)
	options struct {
		resourceExtract func(*gin.Context) string
		blockFallback   func(*gin.Context)
	}
)

func evaluateOptions(opts []Option) *options {
	optCopy := &options{}
	for _, opt := range opts {
		opt(optCopy)
	}

	return optCopy
}

// ResourceExtractor sets the resource extractor of the web requests.
func ResourceExtractor(fn func(*gin.Context) string) Option {
	return func(opts *options) {
		opts.resourceExtract = fn
	}
}

// BlockFallback sets the fallback handler when requests are blocked.
func BlockFallback(fn func(ctx *gin.Context)) Option {
	return func(opts *options) {
		opts.blockFallback = fn
	}
}
