package recovery

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddlewareHandleRecovery(t *testing.T) {
	// Create a new Gin router and set the recovery middleware
	router := gin.New()
	router.Use(gin.CustomRecovery(HandleRecovery))
	// Test case 1: Simulate a panic inside a handler
	router.GET("/panic", func(c *gin.Context) {
		panic("Simulated panic")
	})

	// Test case 2: Simulate a handler without panic
	router.GET("/normal", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Handler without panic"})
	})

	// Perform the test cases
	testCases := []struct {
		url          string
		expectedCode int
	}{
		{"/panic", http.StatusOK},
		{"/normal", http.StatusOK},
	}

	for _, tc := range testCases {
		req, _ := http.NewRequest(http.MethodGet, tc.url, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		if w.Code != tc.expectedCode {
			t.Errorf("Expected status code %d but got %d for URL %s", tc.expectedCode, w.Code, tc.url)
		}
	}
}
