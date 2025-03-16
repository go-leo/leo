package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithInstancerBuilder(consulx.Builder{}),
		grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
	)
	for i := 0; i < 100; i++ {
		r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: randx.HexString(10)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
