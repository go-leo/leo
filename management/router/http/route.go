package http

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	leohttp "github.com/go-leo/leo/v2/http"
)

type HTTPMapping struct {
	HTTPRoutes []HTTPRoute `json:"http_routes,omitempty"`
}

type HTTPRoute struct {
	Path   string `json:"path,omitempty"`
	Method string `json:"method,omitempty"`
}

type HealthCheckOptions struct {
	TLSConf *tls.Config
	Timeout time.Duration
}

func Route(rg *gin.RouterGroup, httpSrv *leohttp.Server, healthCheckOptions *HealthCheckOptions) {
	if httpSrv == nil {
		return
	}
	httpMapping := new(HTTPMapping)
	for _, routeInfo := range httpSrv.Engin().Routes() {
		route := HTTPRoute{
			Path:   routeInfo.Path,
			Method: routeInfo.Method,
		}
		httpMapping.HTTPRoutes = append(httpMapping.HTTPRoutes, route)
	}
	rg.GET("/http/mapping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"http": httpMapping,
		})
	})

	addr := fmt.Sprintf("%s://%s", httpSrv.Scheme(), net.JoinHostPort(httpSrv.Host(), strconv.Itoa(httpSrv.Port())))
	target := path.Join(addr, httpSrv.HealthCheckPath())
	httpProber := NewProber(healthCheckOptions.Timeout, healthCheckOptions.TLSConf)
	rg.GET("/http/health/check", func(c *gin.Context) {
		if err := httpProber.Check(target); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusNoContent)
	})
}
