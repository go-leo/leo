package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"github.com/go-leo/leo/v3/transportx/httpx"
	"github.com/gorilla/mux"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {
	go runApi(8000)
	go runHttp(8080, "red")
	go runHttp(8081, "blue")
	go runHttp(8082, "yellow")
	go runHttp(18080, "red")
	go runHttp(18081, "blue")
	go runHttp(18082, "yellow")
	go runHttp(28080, "red")
	go runHttp(28081, "blue")
	go runHttp(28082, "yellow")
	go runGrpc(9090, "red")
	go runGrpc(9091, "blue")
	go runGrpc(9092, "yellow")
	go runGrpc(19090, "red")
	go runGrpc(19091, "blue")
	go runGrpc(19092, "yellow")
	go runGrpc(29090, "red")
	go runGrpc(29091, "blue")
	go runGrpc(29092, "yellow")
	select {}
}

func runApi(port int) {
	httpClient := helloworld.NewGreeterHttpClient(
		"consul://localhost:8500/demo.http?dc=dc1",
		httpx.InstancerBuilder(consulx.Builder{}),
		httpx.BalancerFactory(lbx.RandomFactory{}),
	)

	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		NewGreeterApiService(httpClient),
	)
	app := leo.NewApp(leo.Runner(httpx.NewServer(router, httpx.Port(port))))
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

type GreeterApiService struct {
	client helloworld.GreeterService
}

func (g GreeterApiService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := g.client.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	hello.Message = "[client," + request.GetName() + "]" + "[api]" + hello.GetMessage()
	return hello, nil
}

func NewGreeterApiService(client helloworld.GreeterService) helloworld.GreeterService {
	return &GreeterApiService{client: client}
}

func runHttp(port int, color string) {
	grpcClient := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/demo.grpc?dc=dc1",
		grpcx.DialOptions(grpc1.WithTransportCredentials(insecure.NewCredentials())),
		grpcx.InstancerBuilder(consulx.Builder{}),
		grpcx.BalancerFactory(lbx.RandomFactory{}),
	)
	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		NewGreeterHttpService(grpcClient, color),
	)
	server := httpx.NewServer(
		router,
		httpx.Port(port),
		httpx.Instance("consul://localhost:8500/demo.http?dc=dc1"),
		httpx.RegistrarBuilder(consulx.Builder{}),
		httpx.Color(color),
	)
	app := leo.NewApp(leo.Runner(server))
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

type GreeterHttpService struct {
	client helloworld.GreeterService
	color  string
}

func (g GreeterHttpService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	hello, err := g.client.SayHello(ctx, request)
	if err != nil {
		return nil, err
	}
	hello.Message = "[http," + g.color + "]" + hello.GetMessage()
	return hello, nil
}

func NewGreeterHttpService(client helloworld.GreeterService, color string) helloworld.GreeterService {
	return &GreeterHttpService{client: client, color: color}
}

func runGrpc(port int, color string) {
	server := grpcx.NewServer(
		grpcx.Port(port),
		grpcx.Instance("consul://localhost:8500/demo.grpc?dc=dc1"),
		grpcx.RegistrarBuilder(consulx.Builder{}),
		grpcx.Color(color),
	)
	service := helloworld.NewGreeterGrpcServer(NewGreeterGrpcService(color))
	helloworld.RegisterGreeterServer(server, service)
	app := leo.NewApp(leo.Runner(server))
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}

type GreeterGrpcService struct {
	color string
}

func (g GreeterGrpcService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	time.Sleep(time.Second)
	return &helloworld.HelloReply{Message: "[grpc," + g.color + "]"}, nil
}

func NewGreeterGrpcService(color string) helloworld.GreeterService {
	return &GreeterGrpcService{color: color}
}
