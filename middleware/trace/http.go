package trace

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func GinMiddleware(serviceName string, opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	return otelgin.Middleware(
		serviceName,
		otelgin.WithPropagators(o.Propagators),
		otelgin.WithTracerProvider(o.TracerProvider),
		otelgin.WithFilter(func(request *http.Request) bool {
			if _, ok := o.Skips[request.URL.Path]; ok {
				return false
			}
			return true
		}),
	)
}
