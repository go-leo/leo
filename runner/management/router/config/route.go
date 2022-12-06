package config

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/global"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/config", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, global.Config())
	})
}
