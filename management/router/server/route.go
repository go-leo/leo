package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GRPCMapping struct {
	FullMethods []string `json:"full_methods,omitempty"`
}

type HTTPMapping struct {
	HTTPRoutes []HTTPRoute `json:"http_routes,omitempty"`
}

type HTTPRoute struct {
	Path   string `json:"path,omitempty"`
	Method string `json:"methods,omitempty"`
}

func Route(rg *gin.RouterGroup, grpcMapping *GRPCMapping, httpMapping *HTTPMapping) {
	rg.GET("/server", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"grpc": grpcMapping,
			"http": httpMapping,
		})
	})
	rg.GET("/server/grpc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"grpc": grpcMapping,
		})
	})
	rg.GET("/server/http", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"http": httpMapping,
		})
	})
}
