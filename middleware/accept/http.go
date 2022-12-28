/*
Package accept contains filters to reject requests without a specified Accept
header with "406 Not Acceptable".
*/
package accept

import (
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	// ALL matches all media types.
	ALL = "*/*"
)

var (
	// EventStream matches media typs used for SSE events.
	EventStream = GinMiddleware("text/event-stream", "text/*")

	// HTML matches media typs used for HTML encoded resources.
	HTML = GinMiddleware("text/html")

	// JSON matches media typs used for JSON encoded resources.
	JSON = GinMiddleware("application/json", "application/javascript")

	// Plain matches media typs used for plaintext resources.
	Plain = GinMiddleware("text/plain")

	// XML matches media typs used for XML encoded resources.
	XML = GinMiddleware("application/xhtml+xml", "application/xml")
)

// GinMiddleware returns a gin.HandlerFunc to restrict accepted
// media types and respond with "406 Not Acceptable" otherwise.
func GinMiddleware(mediaTypes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !acceptable(c.GetHeader("Accept"), mediaTypes) {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
	}
}

func acceptable(accept string, mediaTypes []string) bool {
	if accept == "" || accept == ALL {
		// The absense of an Accept header is equivalent to "*/*".
		// https://tools.ietf.org/html/rfc2296#section-4.2.2
		return true
	}

	for _, a := range strings.Split(accept, ",") {
		mediaType, _, err := mime.ParseMediaType(a)
		if err != nil {
			continue
		}

		if mediaType == ALL {
			return true
		}

		for _, t := range mediaTypes {
			if mediaType == t {
				return true
			}
		}
	}

	return false
}
