package trace

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func Middleware(serviceName string, opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	middleware := otelgin.Middleware(
		serviceName,
		otelgin.WithPropagators(o.Propagators),
		otelgin.WithTracerProvider(o.TracerProvider),
	)
	return func(c *gin.Context) {
		for _, skip := range o.Skips {
			if skip(c) {
				return
			}
		}
		middleware(c)
	}
}
