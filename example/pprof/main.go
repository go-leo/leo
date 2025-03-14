package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/pprofx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// http server
	client := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	router := helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(grpcSrv, helloworld.NewGreeterGrpcServer(&server{}))

	actuatorRouter := mux.NewRouter()
	// 添加pprof路由
	actuatorRouter = pprofx.Append(actuatorRouter)
	actuatorSrv := actuator.NewServer(16060, actuatorRouter)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv, actuatorSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
