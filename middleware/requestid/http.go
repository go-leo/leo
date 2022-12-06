package requestid

import (
	"github.com/gin-gonic/gin"

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/v2/middleware/httpheader"
)

func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		// 1. from context
		requestID, _ = FromContext(c.Request.Context())
		if stringx.IsNotBlank(requestID) {
			return
		}
		// 2. from header
		_, ok := httpheader.FromContext(c.Request.Context())
		if !ok {
			httpheader.GinMiddleware()(c)
		}
		requestID, _ = FromHeader(c.Request.Context())
		if stringx.IsNotBlank(requestID) {
			c.Request = c.Request.WithContext(NewContext(c.Request.Context(), requestID))
			return
		}
		// 3. from trace system traceID
		requestID, _ = FromTrace(c.Request.Context())
		if stringx.IsNotBlank(requestID) {
			c.Request = c.Request.WithContext(NewContext(c.Request.Context(), requestID))
			return
		}
		// 4. generate
		requestID = Generate()
		c.Request = c.Request.WithContext(NewContext(c.Request.Context(), requestID))
	}
}
