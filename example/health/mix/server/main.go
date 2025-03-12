package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx/v2"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/healthx"
	"github.com/go-leo/leo/v3/serverx/actuator"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
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
	// 添加健康检查路由
	actuatorRouter = healthx.Append(actuatorRouter)
	actuatorSrv := actuator.NewServer(16060, actuatorRouter)

	customHealthChecker()

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv, actuatorSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// 自定义健康检查
func customHealthChecker() {
	checker := healthx.NewChecker("custom")
	rand, err := randx.NewChaCha8()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			time.Sleep(time.Duration(rand.Int64N(int64(time.Second))))
			if rand.Int()%2 == 0 {
				checker.Resume()
			} else {
				checker.Shutdown()
			}
		}
	}()
	healthx.RegisterChecker(checker)
}
