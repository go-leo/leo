package requestid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
)

// Middleware return the RequestID middleware.
func Middleware(opts ...Option) gin.HandlerFunc {
	o := &options{}
	o.apply(opts...)
	o.init()

	return func(c *gin.Context) {
		var requestID string

		// 1. from header
		requestID = c.GetHeader(o.HeaderKey)
		if len(requestID) != 0 {
			next(c, o, requestID, false)
			return
		}

		// 2. from TraceContext, TraceContext is a propagator that supports the W3C Trace Context format
		// (https://www.w3.org/TR/trace-context/)
		requestID, _ = fromTrace(c)
		if len(requestID) != 0 {
			next(c, o, requestID, true)
			return
		}

		// 3. generate
		requestID = o.Generator()
		next(c, o, requestID, false)
	}
}

var traceCtxRegExp = regexp.MustCompile("^(?P<version>[0-9a-f]{2})-(?P<traceID>[a-f0-9]{32})-(?P<spanID>[a-f0-9]{16})-(?P<traceFlags>[a-f0-9]{2})(?:-.*)?$")

const (
	supportedVersion  = 0
	traceparentHeader = "traceparent"
)

func fromTrace(c *gin.Context) (string, bool) {
	h := c.GetHeader(traceparentHeader)
	if h == "" {
		return "", false
	}
	matches := traceCtxRegExp.FindStringSubmatch(h)
	if len(matches) == 0 {
		return "", false
	}
	if len(matches[2]) != 32 {
		return "", false
	}
	return matches[2][:32], true
}

func next(c *gin.Context, o *options, requestID string, isTraceContext bool) {
	if o.RewriteTraceContext && !isTraceContext {
		if len(requestID) > 32 {
			requestID = requestID[:32]
		} else if len(requestID) < 32 {
			requestID = requestID + hexString(32-len(requestID))
		}
		h := fmt.Sprintf("%.2x-%s-%s-%s",
			supportedVersion,
			requestID,
			strings.ToLower(hexString(16)),
			"00")
		c.Header(traceparentHeader, h)
	}
	c.Request = c.Request.WithContext(NewContext(c.Request.Context(), requestID))
	o.Handler(c, requestID)
	// Set the id to ensure that the X-Request-ID is in the response
	c.Header(o.HeaderKey, requestID)
	c.Next()
}

func hexString(length int) string {
	b := make([]byte, length)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)[:length]
}
