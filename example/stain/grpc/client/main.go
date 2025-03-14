package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/stainx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand/v2"
	"time"
)

func main() {
	client := helloworld.NewGreeterGrpcClient(
		"consul://localhost:8500/leo.example.sd.grpc?dc=dc1",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithInstancerBuilder(consulx.Builder{}),
		grpctransportx.WithBalancerFactory(lbx.RandomFactory{Seed: time.Now().Unix()}),
	)
	r := rand.New(rand.NewPCG(uint64(time.Now().Unix()), uint64(time.Now().Unix())))
	var colors = []string{"red", "blue", "yellow", "black", "white"}
	for i := 0; i < 100; i++ {
		ctx := context.Background()
		color := colors[r.IntN(len(colors))]
		ctx = stainx.ColorInjector(ctx, color)
		r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: color})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
