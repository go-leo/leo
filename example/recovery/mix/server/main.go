package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/recoveryx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	//mdw := recoveryx.Middleware()
	mdw := recoveryx.Middleware(recoveryx.RecoveryHandler(func(ctx context.Context, p any) (err error) {
		return fmt.Errorf("panic: %v", p)
	}))
	// http server
	// 添加recovery中间件
	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		&server{},
		httptransportx.Middleware(mdw),
	)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(
		grpcSrv,
		helloworld.NewGreeterGrpcServer(
			&server{},
			grpctransportx.Middleware(mdw),
		),
	)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	panic("this is panic")
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
