package actuator

import (
	"sync"
)

var handlers []Handler
var handlersMu sync.RWMutex

func RegisterHandler(handler Handler) {
	handlersMu.Lock()
	defer handlersMu.Unlock()
	handlers = append(handlers, handler)
}

func getHandlers() []Handler {
	handlersMu.RLock()
	result := make([]Handler, len(handlers))
	copy(result, handlers)
	handlersMu.RUnlock()
	return result
}
