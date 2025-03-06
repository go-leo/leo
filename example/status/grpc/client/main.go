package main

import (
	"context"
	"flag"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/example/api/status/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
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
	c := helloworld.NewGreeterGrpcClient(*addr, grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: *name})
	if err != nil {
		st, ok := status.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
