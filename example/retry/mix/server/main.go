package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/retryx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	// http server
	grpcClient := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(
			retryx.Middleware(retryx.MaxAttempts(3), retryx.Backoff(retryx.Exponential())),
		),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		grpcClient,
	)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(
		grpcSrv,
		helloworld.NewGreeterGrpcServer(
			&server{},
		),
	)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	fmt.Println("SayHello", time.Now())
	return nil, statusx.Internal(statusx.RetryInfo(time.Second))
}
