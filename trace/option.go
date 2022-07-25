package trace

import "go.opentelemetry.io/otel/attribute"

type options struct {
	// HTTPOptions
	HTTPOptions *HTTPOptions
	// GRPCOptions
	GRPCOptions *GRPCOptions
	// JaegerOptions
	JaegerOptions *JaegerOptions
	// ZipkinOptions
	ZipkinOptions *ZipkinOptions
	// WriterOptions
	WriterOptions *WriterOptions
	// SampleRate 采样率
	SampleRate float64
	// Attributes trace需要一些额外的信息
	Attributes []attribute.KeyValue
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {

}

type Option func(o *options)

func Jaeger(jaegerOptions *JaegerOptions) Option {
	return func(o *options) {
		o.JaegerOptions = jaegerOptions
	}
}

func Zipkin(zipkinOptions *ZipkinOptions) Option {
	return func(o *options) {
		o.ZipkinOptions = zipkinOptions
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

func SampleRate(rate float64) Option {
	return func(o *options) {
		o.SampleRate = rate
	}
}

func Attributes(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.Attributes = attrs
	}
}
