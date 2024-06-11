package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld"
	consulapi "github.com/hashicorp/consul/api"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	conn, err := grpc1.Dial(":9090", grpc1.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	apiclient, err := consulapi.NewClient(&consulapi.Config{
		Address: "",
	})
	instancer := consul.NewInstancer(consul.NewClient(apiclient), nil, "", nil, true)

	endpoints := helloworld.NewGreeterGrpcClientEndpoints(
		basicx.Middleware("soyacen", "123456", "basic auth example"),
	)
	factories := helloworld.NewGreeterGrpcClientFactories(endpoints)
	retry := lb.Retry(10, time.Second, lb.NewRoundRobin(sd.NewEndpointer(instancer, factories.SayHello(), nil)))

	transports := helloworld.NewGreeterGrpcClientTransports(conn)

	// ok
	client := helloworld.NewGreeterGrpcClient(transports)
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// panic
	client = helloworld.NewGreeterGrpcClient(
		transports,
		basicx.Middleware("soyacen", "654321", "basic auth example"),
	)
	reply, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
