package health

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type HttpOptions struct {
	Target  string
	TLSConf *tls.Config
	Timeout time.Duration
}

type GRPCOptions struct {
	Target  string
	TLSConf *tls.Config
	Timeout time.Duration
}

func Route(rg *gin.RouterGroup, httpOptions *HttpOptions, gRPCOptions *GRPCOptions) {
	httpChecker := httpChecker(httpOptions)
	grpcChecker := grpcChecker(gRPCOptions)
	rg.GET("/health", func(c *gin.Context) {
		if grpcChecker == nil || httpChecker() == nil {
			c.Status(http.StatusNotImplemented)
		}
		eg, _ := errgroup.WithContext(c.Request.Context())
		if grpcChecker != nil {
			eg.Go(func() error {
				return grpcChecker()
			})
		}
		if httpChecker != nil {
			eg.Go(func() error {
				return httpChecker()
			})
		}
		if err := eg.Wait(); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusNoContent)
	})
	rg.GET("/health/grpc", func(c *gin.Context) {
		if grpcChecker == nil {
			c.Status(http.StatusNotImplemented)
		}
		if err := grpcChecker(); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusNoContent)
	})
	rg.GET("/health/http", func(c *gin.Context) {
		if httpChecker == nil {
			c.Status(http.StatusNotImplemented)
		}
		if err := httpChecker(); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusNoContent)
	})
}

func httpChecker(httpConf *HttpOptions) func() error {
	if httpConf == nil {
		return nil
	}
	httpProber := NewHTTPProber(httpConf.Timeout, httpConf.TLSConf)
	return func() error {
		return httpProber.Check(httpConf.Target)
	}
}

func grpcChecker(grpcConf *GRPCOptions) func() error {
	if grpcConf == nil {
		return nil
	}
	grpcProber := NewGRPCProber(grpcConf.Timeout, grpcConf.TLSConf)
	return func() error {
		return grpcProber.Check(grpcConf.Target)
	}
}
