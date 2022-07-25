package metric

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/global"
)

func Route(rg *gin.RouterGroup) {
	prometheusExporter := global.GetPrometheusExporter()
	if prometheusExporter != nil {
		rg.GET("/metrics", gin.WrapH(prometheusExporter))
	}
}
