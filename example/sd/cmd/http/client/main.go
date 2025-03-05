package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/sd/api"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
)

func main() {
	client := api.NewGreeterHttpClient(
		"consul://localhost:8500/leo.example.sd.http?dc=dc1",
		httptransportx.WithInstancerBuilder(consulx.Builder{}),
		httptransportx.WithBalancerFactory(lbx.RoundRobinFactory{}),
	)
	for i := 0; i < 10; i++ {
		r, err := client.SayHello(context.Background(), &api.HelloRequest{Name: randx.HexString(10)})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
