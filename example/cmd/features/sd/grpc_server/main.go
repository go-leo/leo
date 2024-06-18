package main

import (
	"context"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	"github.com/go-leo/leo/v3/logx"
	"github.com/google/uuid"
	stdconsul "github.com/hashicorp/consul/api"
	grpc1 "google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	go run(9090)
	go run(9091)
	go run(9092)
	select {}
}

func run(port int) {
	address := ":" + strconv.FormatInt(int64(port), 10)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc1.NewServer()
	endpoints := helloworld.NewGreeterServerEndpoints(NewGreeterService(address))
	transports := helloworld.NewGreeterGrpcServerTransports(endpoints)
	service := helloworld.NewGreeterGrpcServer(transports)
	helloworld.RegisterGreeterServer(s, service)
	client, err := stdconsul.NewClient(&stdconsul.Config{
		Address:    "localhost:8500",
		Datacenter: "dc1",
	})
	if err != nil {
		panic(err)
	}
	registrar := consul.NewRegistrar(consul.NewClient(client), &stdconsul.AgentServiceRegistration{
		ID:      uuid.NewString(),
		Name:    "demo.grpc",
		Port:    port,
		Address: "127.0.0.1",
	}, logx.L())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		registrar.Deregister()
		s.GracefulStop()
	}()

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		registrar.Register()
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v\n", err)
		}
	}()
	select {}
}

type GreeterService struct {
	address string
}

func (g GreeterService) SayHello(ctx context.Context, request *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	time.Sleep(time.Millisecond)
	return &helloworld.HelloReply{Message: "hi " + request.GetName() + ", i am " + g.address}, nil
}

func NewGreeterService(address string) helloworld.GreeterService {
	return &GreeterService{address: address}
}
