package requestid

import (
	"encoding/hex"
	"math/rand"
	"sync"
	"time"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/stringx"
	"github.com/gin-gonic/gin"
)

type options struct {
	Generator           func() string
	HeaderKey           string
	Handler             func(c *gin.Context, requestID string)
	RewriteTraceContext bool
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Generator == nil {
		o.Generator = func(randPool *sync.Pool) func() string {
			return func() string {
				randSource := randPool.Get().(*rand.Rand)
				defer randPool.Put(randSource)
				var tid [16]byte
				randSource.Read(tid[:])
				return hex.EncodeToString(tid[:])
			}
		}(&sync.Pool{New: func() any { return rand.New(rand.NewSource(time.Now().UnixNano())) }})
	}
	if stringx.IsBlank(o.HeaderKey) {
		o.HeaderKey = "X-Request-ID"
	}
	if o.Handler == nil {
		o.Handler = func(_ *gin.Context, _ string) {}
	}
}

// Option for queue system
type Option func(*options)

// IDGenerator set id Generator function
func IDGenerator(g func() string) Option {
	return func(cfg *options) {
		cfg.Generator = g
	}
}

// CustomHeaderKey set custom header key for request id
func CustomHeaderKey(key string) Option {
	return func(cfg *options) {
		cfg.HeaderKey = key
	}
}

// Handler set Handler function for request id with context
func Handler(handler func(c *gin.Context, requestID string)) Option {
	return func(cfg *options) {
		cfg.Handler = handler
	}
}

// RewriteTraceContext rewrite trace context if trace context not exist
func RewriteTraceContext() Option {
	return func(o *options) {
		o.RewriteTraceContext = true
	}
}
