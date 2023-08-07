package trace

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/propagation"
)

func TestMiddleware(t *testing.T) {
	r := gin.New()
	r.Use(Middleware("test-service", Propagators(propagation.TraceContext{})))

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "Test route")
	})

	// Example:
	resp := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	r.ServeHTTP(resp, req)
	if resp.Code != 200 {
		t.Errorf("Expected status code 200, got %d", resp.Code)
	}
}
