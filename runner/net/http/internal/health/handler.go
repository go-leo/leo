package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerFunc(srv *Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, srv.Check(c.Request.Context()))
	}
}
