package actuator

import (
	"net/http"
)

// Handler is a actuator handler
type Handler interface {
	// Pattern return the handler path
	Pattern() string
	// Handle handle the request and response to caller
	Handle(http.ResponseWriter, *http.Request)
}

// HandlerProvider is a handler provider
type HandlerProvider interface {
	// ActuatorHandler return a handler
	ActuatorHandler() Handler
}
