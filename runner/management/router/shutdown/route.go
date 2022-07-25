package shutdown

import (
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Route(rg *gin.RouterGroup, signals []os.Signal) {
	rg.GET("/shutdown", func(c *gin.Context) {
		for _, signal := range signals {
			signum, ok := signal.(syscall.Signal)
			if ok {
				if err := syscall.Kill(syscall.Getpid(), signum); err == nil {
					c.Status(http.StatusOK)
					return
				}
			}
		}
		c.JSON(http.StatusInternalServerError, gin.H{"signals": signals})
	})
}
