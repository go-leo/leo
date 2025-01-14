package metrics

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.27.0"

	"github.com/go-leo/leo/middleware/internal"
	"github.com/go-leo/leo/middleware/noop"
)

func GinMiddleware(opts ...Option) gin.HandlerFunc {
	o := new(options)
	o.apply(opts...)
	o.init()
	skipMap := make(map[string]struct{}, len(o.Skips))
	for _, skip := range o.Skips {
		skipMap[skip] = struct{}{}
	}
	// 请求延迟直方图
	latencyHistogram, e := otel.GetMeterProvider().
		Meter(internal.InstrumentationName).
		Float64Histogram(
			"http.server.latency",
			metric.WithDescription("The HTTP request latencies in seconds."),
		)
	if e != nil {
		otel.Handle(e)
		return noop.GinMiddleware()
	}
	// 请求计数器
	requestCounter, err := otel.GetMeterProvider().
		Meter(internal.InstrumentationName).
		Int64Counter(
			"http.server.requests",
			metric.WithDescription("How many HTTP requests processed, partitioned by status code and HTTP method."),
		)
	if err != nil {
		otel.Handle(err)
		return noop.GinMiddleware()
	}
	// 请求大小
	requestSizeCounter, err := otel.GetMeterProvider().
		Meter(internal.InstrumentationName).
		Int64UpDownCounter("http.server.request.size.bytes",
			metric.WithDescription("The HTTP request sizes in bytes."),
		)
	if err != nil {
		otel.Handle(err)
		return noop.GinMiddleware()
	}
	// 响应大小
	responseSizeCounter, err := otel.GetMeterProvider().
		Meter(internal.InstrumentationName).
		Int64UpDownCounter("http.server.response.size.bytes",
			metric.WithDescription("The HTTP response sizes in bytes."),
		)
	if err != nil {
		otel.Handle(err)
		return noop.GinMiddleware()
	}
	return func(c *gin.Context) {
		if _, ok := skipMap[c.FullPath()]; ok {
			c.Next()
			return
		}
		// 开始时间
		startTime := time.Now()

		// 记录请求大小
		reqSz := computeApproximateRequestSize(c.Request)
		requestSizeCounter.Add(c.Request.Context(), int64(reqSz))

		// 处理中间件和业务逻辑
		c.Next()

		// 记录响应大小
		responseSizeCounter.Add(c.Request.Context(), int64(c.Writer.Size()))

		// 包含接口信息的属性
		attrs := []attribute.KeyValue{
			internal.RPCSystemHTTPServer,
			semconv.HTTPRouteKey.String(c.FullPath()),
			semconv.HTTPRequestMethodKey.String(c.Request.Method),
			semconv.HTTPResponseStatusCodeKey.Int(c.Writer.Status()),
		}

		// 请求延迟直方图记录延迟
		latencyHistogram.Record(c.Request.Context(), time.Since(startTime).Seconds(), metric.WithAttributeSet(attribute.NewSet(attrs...)))

		// 请求计数器加1
		requestCounter.Add(c.Request.Context(), 1, metric.WithAttributeSet(attribute.NewSet(attrs...)))
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
