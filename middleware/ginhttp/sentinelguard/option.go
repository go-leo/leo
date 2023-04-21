package sentinelguard

import (
	"github.com/gin-gonic/gin"
)

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
