package trace

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/go-leo/leo/runner/net/http/client"
)

const instrumentationName = "github.com/go-leo/leo"

func HTTPClientMiddleware(opts ...Option) client.Interceptor {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	return func(ctx context.Context, req any, reply any, info *client.HTTPInfo, invoke client.Invoker) error {
		if _, ok := skipMap[info.Path]; ok {
			return invoke(ctx, req, reply, info)
		}
		var tracer trace.Tracer
		if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
			tracer = span.TracerProvider().Tracer(instrumentationName, trace.WithInstrumentationVersion(otelhttp.SemVersion()))
		} else {
			tracer = o.TracerProvider.Tracer(instrumentationName, trace.WithInstrumentationVersion(otelhttp.SemVersion()))
		}
		opts := append([]trace.SpanStartOption{}, trace.WithSpanKind(trace.SpanKindClient)) // start with the configured options

		spanName := "HTTP " + info.Request.URL.Path
		ctx, span := tracer.Start(ctx, spanName, opts...)
		defer span.End()

		info.Request = info.Request.WithContext(ctx)
		span.SetAttributes(semconv.HTTPClientAttributesFromHTTPRequest(info.Request)...)

		o.Propagators.Inject(ctx, propagation.HeaderCarrier(info.Request.Header))

		err := invoke(ctx, req, reply, info)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return err
		}

		span.SetAttributes(semconv.HTTPAttributesFromHTTPStatusCode(info.Response.StatusCode)...)
		span.SetStatus(semconv.SpanStatusFromHTTPStatusCode(info.Response.StatusCode))
		return err
	}
}

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
