package main

import (
	"context"
	"github.com/RussellLuo/slidingwindow"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/ratelimitx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"log"
	"time"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	// 滑动窗口，每秒10个请求
	limiter, stopFunc := slidingwindow.NewLimiter(time.Second, 10, func() (slidingwindow.Window, slidingwindow.StopFunc) { return slidingwindow.NewLocalWindow() })
	// 退出时关闭
	defer stopFunc()
	mdw := ratelimitx.SlideWindow(limiter)
	grpcServer := helloworld.NewGreeterGrpcServer(
		&server{},
		grpctransportx.Middleware(mdw),
	)
	helloworld.RegisterGreeterServer(grpcSrv, grpcServer)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
