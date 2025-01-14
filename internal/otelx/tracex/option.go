package tracex

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-leo/leo/internal/otelx/resourcex"
)

type options struct {
	// Service 服务
	Service *resourcex.Service
	// Attributes trace需要一些额外的信息
	Attributes []attribute.KeyValue
	// Resources 资源
	Resources resourcex.ResourceFlag
	// HTTPOptions
	HTTPOptions *HTTPOptions
	// GRPCOptions
	GRPCOptions *GRPCOptions
	// JaegerOptions
	JaegerOptions *JaegerOptions
	// WriterOptions
	WriterOptions *WriterOptions
	// Sampler 自定义Sampler
	Sampler sdktrace.Sampler
	// IDGenerator 自定义id生成器
	IDGenerator sdktrace.IDGenerator
	// SpanProcessor 自定义span处理器
	SpanProcessor sdktrace.SpanProcessor
	// RawSpanLimits
	RawSpanLimits *sdktrace.SpanLimits
	Propagators   []propagation.TextMapPropagator
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.Propagators == nil {
		o.Propagators = []propagation.TextMapPropagator{
			propagation.Baggage{},
			propagation.TraceContext{},
		}
	}
}

type Option func(o *options)

func Service(svc *resourcex.Service) Option {
	return func(o *options) {
		o.Service = svc
	}
}

func Attributes(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.Attributes = attrs
	}
}

func Resources(res resourcex.ResourceFlag) Option {
	return func(o *options) {
		o.Resources = res
	}
}

func Jaeger(jaegerOptions *JaegerOptions) Option {
	return func(o *options) {
		o.JaegerOptions = jaegerOptions
	}
}

func Writer(writerOptions *WriterOptions) Option {
	return func(o *options) {
		o.WriterOptions = writerOptions
	}
}

func GRPC(gRPCOptions *GRPCOptions) Option {
	return func(o *options) {
		o.GRPCOptions = gRPCOptions
	}
}

func HTTP(httpOptions *HTTPOptions) Option {
	return func(o *options) {
		o.HTTPOptions = httpOptions
	}
}

func Sampler(sampler sdktrace.Sampler) Option {
	return func(o *options) {
		o.Sampler = sampler
	}
}

func IDGenerator(custom bool, idGen sdktrace.IDGenerator) Option {
	return func(o *options) {
		if custom {
			o.IDGenerator = idGen
		}
	}
}

func SpanProcessor(spanProcessor sdktrace.SpanProcessor) Option {
	return func(o *options) {
		o.SpanProcessor = spanProcessor
	}
}

func RawSpanLimits(limits *sdktrace.SpanLimits) Option {
	return func(o *options) {
		o.RawSpanLimits = limits
	}
}

func Propagators(propagators ...propagation.TextMapPropagator) Option {
	return func(o *options) {
		o.Propagators = append(o.Propagators, propagators...)
	}
}
