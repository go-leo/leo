package actuator

import (
	"net/http"
)

type Handler interface {
	Pattern() string
	Handle(http.ResponseWriter, *http.Request)
}

type HandlerProvider interface {
	ActuatorHandler() Handler
}
