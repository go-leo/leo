package sentinelguard

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	// Test cases

	t.Run("Request allowed", func(t *testing.T) {
		// Create a new Gin context for testing
		router := gin.New()
		router.Use(Middleware())
		router.GET("/api/users/1", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		// Create a request to test
		req, _ := http.NewRequest("GET", "/api/users/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response status code is 200 (OK) as the request is allowed
		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, rec.Code)
		}
	})

	t.Run("Request blocked", func(t *testing.T) {
		// Create a new Gin context for testing with a custom block fallback
		router := gin.New()
		router.Use(Middleware(BlockFallback(func(c *gin.Context) {
			c.AbortWithStatus(http.StatusForbidden)
		})))
		router.GET("/api/users/1", func(c *gin.Context) {
			c.JSON(http.StatusTooManyRequests, "ok")
		})

		// Create a request to test
		req, _ := http.NewRequest("GET", "/api/users/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response status code is 429 (Too Many Requests) as the request is blocked
		if rec.Code != http.StatusTooManyRequests {
			t.Errorf("Expected status code %d but got %d", http.StatusTooManyRequests, rec.Code)
		}
	})

	t.Run("Request Sentry", func(t *testing.T) {
		// Create a new Gin context for testing with a custom block fallback
		router := gin.New()
		router.Use(Middleware(ResourceExtractor(func(c *gin.Context) string {
			c.AbortWithStatus(http.StatusBadRequest)
			return "test"
		})))
		router.GET("/api/users/1", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ok")
		})

		// Create a request to test
		req, _ := http.NewRequest("GET", "/api/users/1", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		// Assert the response status code is 429 (Too Many Requests) as the request is blocked
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d but got %d", http.StatusTooManyRequests, rec.Code)
		}

	})

	// Add more test cases to cover different scenarios as needed.
}
