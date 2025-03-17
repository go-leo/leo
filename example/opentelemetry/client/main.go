package main

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"net/http"
	"os"
)

func init() {
	// init trace
	file, err := os.OpenFile("/tmp/leo.client.trace.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	stdoutExporter, err := stdout.New(stdout.WithWriter(file), stdout.WithPrettyPrint())
	if err != nil {
		panic(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(stdoutExporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
}

func main() {
	shutdown := otel.GetTracerProvider().(interface {
		Shutdown(ctx context.Context) error
	})
	defer shutdown.Shutdown(context.Background())

	// 创建 otel 的 http Transport, 包含了 trace
	httpCli := httptransport.SetClient(&http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)})
	client := helloworld.NewGreeterHttpClient(
		"localhost:60051",
		httptransportx.WithClientTransportOption(httpCli),
	)
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: randx.HexString(10)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
