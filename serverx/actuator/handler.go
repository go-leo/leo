package actuator

//
//import (
//	"golang.org/x/exp/slices"
//	"net/http"
//	"sync"
//)
//
//var (
//	handlers   []Handler
//	handlersMu sync.RWMutex
//)
//
//func RegisterHandler(handler Handler) {
//	handlersMu.Lock()
//	defer handlersMu.Unlock()
//	handlers = append(handlers, handler)
//}
//
//func GetHandlers() []Handler {
//	handlersMu.RLock()
//	result := slices.Clone(handlers)
//	handlersMu.RUnlock()
//	return result
//}
//
//// Handler is a actuator handler
//type Handler interface {
//	// Pattern return the handler path
//	Pattern() string
//	// Handler handle the request
//	http.Handler
//}
//
//// HandlerProvider is a handler provider
//type HandlerProvider interface {
//	// ActuatorHandler return a handler
//	ActuatorHandler() Handler
//}
