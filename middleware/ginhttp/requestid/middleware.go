package requestid

import (
	"github.com/gin-gonic/gin"
	"github.com/go-leo/gox/stringx"

	"codeup.aliyun.com/qimao/leo/leo/pkg/requestid"
)

// New initializes the RequestID middleware.
func New(opts ...Option) gin.HandlerFunc {
	o := &options{}
	o.apply(opts...)
	o.init()

	return func(c *gin.Context) {
		var requestID string

		// 1. from context
		requestID, _ = requestid.FromContext(c.Request.Context())
		if stringx.IsNotBlank(requestID) {
			next(c, o, requestID)
			return
		}

		// 2. from header
		requestID = c.GetHeader(o.headerKey)
		if stringx.IsNotBlank(requestID) {
			c.Request = c.Request.WithContext(requestid.NewContext(c.Request.Context(), requestID))
			next(c, o, requestID)
			return
		}

		// 3. generate
		requestID = o.generator()
		c.Request.Header.Add(o.headerKey, requestID)
		c.Request = c.Request.WithContext(requestid.NewContext(c.Request.Context(), requestID))
		next(c, o, requestID)
		next(c, o, requestID)
	}
}

func next(c *gin.Context, o *options, requestID string) {
	o.handler(c, requestID)
	// Set the id to ensure that the X-Request-ID is in the response
	c.Header(o.headerKey, requestID)
	c.Next()
}
