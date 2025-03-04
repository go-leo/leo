package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/timeout/api"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	c := api.NewGreeterGrpcClient(*addr, grpcx.DialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
