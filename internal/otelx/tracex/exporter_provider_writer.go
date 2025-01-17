package tracex

import (
	"context"
	"io"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type WriterOptions struct {
	// Writer 标准输入或者文件
	Writer            io.Writer
	PrettyPrint       bool
	WithoutTimestamps bool
}

func (o *WriterOptions) Exporter(ctx context.Context) (sdktrace.SpanExporter, error) {
	opts := []stdouttrace.Option{
		stdouttrace.WithWriter(o.Writer),
	}
	if o.PrettyPrint {
		opts = append(opts, stdouttrace.WithPrettyPrint())
	}
	if o.WithoutTimestamps {
		opts = append(opts, stdouttrace.WithoutTimestamps())
	}
	return stdouttrace.New(opts...)
}
