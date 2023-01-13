package log

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/leo/v2/global"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/log/level", gin.WrapH(global.Logger()))
	rg.PUT("/log/level", gin.WrapH(global.Logger()))
}
