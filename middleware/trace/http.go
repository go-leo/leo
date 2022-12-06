package trace

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func GinMiddleware(serviceName string, opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	middleware := otelgin.Middleware(serviceName, otelgin.WithPropagators(o.Propagators), otelgin.WithTracerProvider(o.TracerProvider))
	return func(c *gin.Context) {
		if _, ok := skipMap[c.FullPath()]; ok {
			c.Next()
			return
		}
		middleware(c)
	}
}
