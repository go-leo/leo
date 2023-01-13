package profile

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup) {
	router := rg.Group("/debug/pprof")
	router.Any("/", gin.WrapF(pprof.Index))
	router.GET("/cmdline", gin.WrapF(pprof.Cmdline))
	router.GET("/profile", gin.WrapF(pprof.Profile))
	router.POST("/symbol", gin.WrapF(pprof.Symbol))
	router.GET("/symbol", gin.WrapF(pprof.Symbol))
	router.GET("/trace", gin.WrapF(pprof.Trace))
	router.GET("/allocs", gin.WrapF(pprof.Handler("allocs").ServeHTTP))
	router.GET("/block", gin.WrapF(pprof.Handler("block").ServeHTTP))
	router.GET("/goroutine", gin.WrapF(pprof.Handler("goroutine").ServeHTTP))
	router.GET("/heap", gin.WrapF(pprof.Handler("heap").ServeHTTP))
	router.GET("/mutex", gin.WrapF(pprof.Handler("mutex").ServeHTTP))
	router.GET("/threadcreate", gin.WrapF(pprof.Handler("threadcreate").ServeHTTP))
}
