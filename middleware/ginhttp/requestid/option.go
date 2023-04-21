package requestid

import (
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/gox/stringx"
)

var randSource = rand.New(rand.NewSource(time.Now().UnixNano()))

type options struct {
	generator func() string
	headerKey string
	handler   func(c *gin.Context, requestID string)
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.generator == nil {
		o.generator = func() string {
			var tid [16]byte
			randSource.Read(tid[:])
			requestID := hex.EncodeToString(tid[:])
			return requestID
		}
	}
	if stringx.IsBlank(o.headerKey) {
		o.headerKey = "X-Request-ID"
	}
	if o.handler == nil {
		o.handler = func(_ *gin.Context, _ string) {}
	}

}

// Option for queue system
type Option func(*options)

// IDGenerator set id generator function
func IDGenerator(g func() string) Option {
	return func(cfg *options) {
		cfg.generator = g
	}
}

// CustomHeaderKey set custom header key for request id
func CustomHeaderKey(key string) Option {
	return func(cfg *options) {
		cfg.headerKey = key
	}
}

// Handler set handler function for request id with context
func Handler(handler func(c *gin.Context, requestID string)) Option {
	return func(cfg *options) {
		cfg.handler = handler
	}
}
