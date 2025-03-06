package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/example/api/status/v1"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 60051))
		if err != nil {
			return err
		}
		client := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
		s := &http.Server{
			Handler: helloworld.AppendGreeterHttpServerRoutes(mux.NewRouter(), client),
		}
		log.Printf("http server listening at %v", lis.Addr())
		return s.Serve(lis)
	})
	eg.Go(func() error {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		helloworld.RegisterGreeterServer(s, helloworld.NewGreeterGrpcServer(&server{}))
		log.Printf("grpc server listening at %v", lis.Addr())
		return s.Serve(lis)
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(_ context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	if in.GetName() == "" {
		// 返回错误
		return nil, status.ErrInvalidName(statusx.RequestInfo(uuid.NewString(), in.GetName()))
	}
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
