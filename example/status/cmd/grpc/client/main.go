package main

import (
	"context"
	"flag"
	errorspb "github.com/go-leo/leo/v3/example/status/api/errors"
	helloworldpb "github.com/go-leo/leo/v3/example/status/api/helloworld"
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
	c := helloworldpb.NewGreeterGrpcClient(*addr, grpcx.DialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	r, err := c.SayHello(context.Background(), &helloworldpb.HelloRequest{Name: *name})
	if err != nil {
		st, ok := errorspb.IsInvalidName(err)
		if ok {
			log.Fatalf("could not greet: %v, identifier: %v, request info: %v", err, st.Identifier(), st.RequestInfo())
		}
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
