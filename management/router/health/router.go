package health

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/v2/management/router/health/internal"
)

type HttpOptions struct {
	Path    string
	Target  string
	TLSConf *tls.Config
	Timeout time.Duration
}

type GRPCOptions struct {
	Path    string
	Target  string
	TLSConf *tls.Config
	Timeout time.Duration
}

type options struct {
	HttpOptions *HttpOptions
	GRPCOptions *GRPCOptions
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.HttpOptions != nil && stringx.IsNotBlank(o.HttpOptions.Path) {
		o.HttpOptions.Path = "/health/http"
	}
	if o.GRPCOptions != nil && stringx.IsNotBlank(o.GRPCOptions.Path) {
		o.GRPCOptions.Path = "/health/grpc"
	}
}

type Option func(o *options)

func Route(rg gin.IRoutes, opts ...Option) {
	o := new(options)
	o.apply(opts...)
	o.init()
	if o.HttpOptions != nil {
		httpProber := internal.NewHTTPProber(o.HttpOptions.Timeout, o.HttpOptions.TLSConf)
		rg.GET(o.HttpOptions.Path, func(c *gin.Context) {
			if err := httpProber.Check(o.HttpOptions.Target); err != nil {
				c.Status(http.StatusServiceUnavailable)
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
	if o.GRPCOptions != nil {
		grpcProber := internal.NewGRPCProber(o.GRPCOptions.Timeout, o.GRPCOptions.TLSConf)
		rg.GET(o.GRPCOptions.Path, func(c *gin.Context) {
			if err := grpcProber.Check(o.GRPCOptions.Target); err != nil {
				c.Status(http.StatusServiceUnavailable)
				return
			}
			c.Status(http.StatusNoContent)
		})
	}
}
