package grpc

import (
	"crypto/tls"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	leogrpc "github.com/go-leo/leo/v2/grpc"
)

type MethodDesc struct {
	MethodName string `json:"method_name,omitempty"`
}

type StreamDesc struct {
	StreamName string `json:"stream_name,omitempty"`
}

type GRPCMapping struct {
	ServiceName string       `json:"service_name,omitempty"`
	Methods     []MethodDesc `json:"methods,omitempty"`
	Streams     []StreamDesc `json:"streams,omitempty"`
}

type HealthCheckOptions struct {
	TLSConf *tls.Config
	Timeout time.Duration
}

func Route(rg *gin.RouterGroup, grpcSrv *leogrpc.Server, healthCheckOptions *HealthCheckOptions) {
	services := grpcSrv.Services()
	grpcMapping := make([]GRPCMapping, 0, len(services))
	for _, service := range services {
		mapping := GRPCMapping{
			ServiceName: service.Desc.ServiceName,
		}
		for _, method := range service.Desc.Methods {
			mapping.Methods = append(mapping.Methods, MethodDesc{MethodName: method.MethodName})
		}
		for _, stream := range service.Desc.Streams {
			mapping.Methods = append(mapping.Methods, MethodDesc{MethodName: stream.StreamName})
		}
		grpcMapping = append(grpcMapping, mapping)
	}
	rg.GET("/grpc/mapping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"grpc": grpcMapping,
		})
	})
	target := net.JoinHostPort(grpcSrv.Host(), strconv.Itoa(grpcSrv.Port()))
	grpcProber := NewProber(healthCheckOptions.Timeout, healthCheckOptions.TLSConf)
	rg.GET("/grpc/health/check", func(c *gin.Context) {
		if err := grpcProber.Check(target); err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusNoContent)
	})
}
