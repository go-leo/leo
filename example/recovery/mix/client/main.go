package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/statusx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	grpcCli := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	Call(grpcCli)
	httpCli := helloworld.NewGreeterHttpClient("localhost:60051")
	Call(httpCli)
}

func Call(grpcCli helloworld.GreeterService) {
	r, err := grpcCli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "recovery"})
	if err != nil {
		st, ok := statusx.From(err)
		if ok {
			log.Printf("could not greet: %v, debugInfo: %v", err, st.DebugInfo())
		}
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
