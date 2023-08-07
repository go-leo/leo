package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMiddleware(t *testing.T) {
	// Initialize a new Gin router and set up the request ID middleware
	r := gin.New()
	r.Use(Middleware())

	// Set up a test route to handle the requests
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test Route",
		})
	})

	// Create a new HTTP request for the "/test" route
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("X-Request-ID", "xxxxxxxxxxxxxxxxxxxxxx")
	// Create a new HTTP recorder to capture the response
	recorder := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Check if the response has the "X-Request-ID" header
	if requestId := recorder.Header().Get("X-Request-ID"); requestId == "" {
		t.Errorf("Expected X-Request-ID header in the response, but got an empty string.")
	}

	// Check if the response JSON contains the expected message
	expectedResponse := `{"message":"Test Route"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, recorder.Body.String())
	}
}

func TestCustomHeaderKeyOption(t *testing.T) {
	// Initialize a new Gin router and set up the request ID middleware with a custom header key
	r := gin.New()
	customHeaderKey := "Custom-Request-ID"
	r.Use(Middleware(CustomHeaderKey(customHeaderKey)))

	// Set up a test route to handle the requests
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Test Route",
		})
	})

	// Create a new HTTP request for the "/test" route
	req, _ := http.NewRequest("GET", "/test", nil)
	// Create a new HTTP recorder to capture the response
	recorder := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(recorder, req)

	// Check if the response has the custom header key
	if requestId := recorder.Header().Get(customHeaderKey); requestId == "" {
		t.Errorf("Expected %s header in the response, but got an empty string.", customHeaderKey)
	}

	// Check if the response JSON contains the expected message
	expectedResponse := `{"message":"Test Route"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response body %s, but got %s", expectedResponse, recorder.Body.String())
	}
}

// Add more tests to cover other scenarios and options if needed.
