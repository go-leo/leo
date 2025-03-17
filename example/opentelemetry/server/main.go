package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/prometheusx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
)

func init() {
	// init metrics
	prometheusExporter, err := prometheus.New()
	if err != nil {
		log.Fatal(err)
	}
	provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(prometheusExporter))
	otel.SetMeterProvider(provider)

	// init trace
	file, err := os.OpenFile("/tmp/leo.server.trace.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
	// http server

	// 添加 open-telemetry grpc 客户端组件, 包含了metrics和trace
	grpcCliOtelHandler := otelgrpc.NewClientHandler(
		otelgrpc.WithMetricAttributes(attribute.String("transport", grpctransportx.GrpcClient)),
		otelgrpc.WithSpanAttributes(attribute.String("transport", grpctransportx.GrpcClient)),
	)
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithDialOptions(
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(grpcCliOtelHandler),
		))

	router := mux.NewRouter()
	// 添加 open-telemetry mux 组件, 包含了metrics和trace
	httpSrvOtelMdw := otelmux.Middleware(
		"leo.example.opentelemetry.gorilla",
		otelmux.WithMetricAttributesFn(func(r *http.Request) []attribute.KeyValue {
			return []attribute.KeyValue{attribute.String("transport", httptransportx.HttpServer)}
		}),
	)
	router.Use(httpSrvOtelMdw)
	router = helloworld.AppendGreeterHttpServerRoutes(router, client)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrvOtelHandler := otelgrpc.NewServerHandler(
		otelgrpc.WithMetricAttributes(attribute.String("transport", grpctransportx.GrpcServer)),
		otelgrpc.WithSpanAttributes(attribute.String("transport", grpctransportx.GrpcServer)),
	)
	grpcSrv := grpcserverx.NewServer(
		grpcserverx.Port(50051),
		// 添加 open-telemetry grpc 服务端组件, 包含了metrics和trace
		grpcserverx.ServerOptions(grpc.StatsHandler(grpcSrvOtelHandler)),
	)
	counter, err := otel.Meter("GreeterServer-Meter").Int64Counter(
		"Call_SayHello",
		metric.WithDescription("call times."),
	)
	if err != nil {
		panic(err)
	}
	greeterService := &server{
		tracer:  otel.Tracer("GreeterServer-Tracer"),
		counter: counter,
	}
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(greeterService))

	actuatorRouter := mux.NewRouter()
	// 添加 prometheus 路由
	actuatorRouter = prometheusx.Append(actuatorRouter)
	actuatorSrv := actuator.NewServer(16060, actuatorRouter)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv, actuatorSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct {
	tracer  oteltrace.Tracer
	counter metric.Int64Counter
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 添加一个自定义的 trace span
	_, span := s.tracer.Start(ctx, "Call_SayHello", oteltrace.WithAttributes(attribute.String("name", in.GetName())))
	defer span.End()
	// 添加一个自定义的 metrics
	s.counter.Add(ctx, 1)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
