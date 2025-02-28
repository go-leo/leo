package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/status/api"
	"github.com/go-leo/leo/v3/transportx/grpcx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "", "Name to greet")
)

func main() {
	flag.Parse()
	c := api.NewGreeterGrpcClient(*addr, grpcx.DialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	r, err := c.SayHello(context.Background(), &api.HelloRequest{Name: *name})
	if err != nil {
		st, ok := api.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
