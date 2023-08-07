package ratelimiter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	// Create a mock context with a request and a mock limiter
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest("GET", "/", nil)
	limiter := alwaysLimiter{v: true}

	// Create the middleware function
	mw := Middleware(limiter)

	// Call the middleware function
	mw(ctx)

	// Check that the response status is 429 Too Many Requests
	assert.Equal(t, http.StatusTooManyRequests, ctx.Writer.Status())
}

func TestAlwaysFalseLimiter(t *testing.T) {
	// Create a mock context with a request
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	ctx.Request, _ = http.NewRequest("GET", "/", nil)

	// Check that the limiter always returns false
	limiter := alwaysFalseLimiter
	assert.False(t, limiter.Limit(ctx))
}
