package log

import (
	"github.com/gin-gonic/gin"

	"github.com/hmldd/leo/global"
)

func Route(rg *gin.RouterGroup) {
	rg.GET("/log/level", gin.WrapH(global.Logger()))
	rg.PUT("/log/level", gin.WrapH(global.Logger()))
}
