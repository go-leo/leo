package clientip

import (
	"github.com/gin-gonic/gin"
	"github.com/go-leo/netx/addrx"
	"github.com/go-leo/stringx"
)

type options struct {
	ToResponse        bool
	ResponseHeaderKey string
}

type Option func(o *options)

func (o *options) init() {
	if o.ToResponse && stringx.IsBlank(o.ResponseHeaderKey) {
		o.ResponseHeaderKey = "X-Real-Client-Ip"
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func ToResponse(toResponse bool) Option {
	return func(o *options) {
		o.ToResponse = toResponse
	}
}

func ResponseHeaderKey(respHeaderKey string) Option {
	return func(o *options) {
		o.ResponseHeaderKey = respHeaderKey
	}
}

func GinMiddleware(opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	return func(c *gin.Context) {
		r := c.Request
		ip := addrx.ClientPublicIP(r)
		if ip == "" {
			ip = addrx.ClientIP(r)
		}
		ctx := r.Context()
		ctx = NewContext(ctx, ip)
		c.Request = r.WithContext(ctx)
		c.Next()
		if o.ToResponse {
			c.Header(o.ResponseHeaderKey, ip)
		}
	}
}
