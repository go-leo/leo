package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/sdx/consulx"
	"github.com/go-leo/leo/v3/sdx/lbx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
)

func main() {
	client := cqrs.NewCqrsHttpClient(
		"consul://localhost:8500/leo.example.cqrs.http?dc=dc1",
		httptransportx.WithInstancerBuilder(consulx.Builder{}),
		httptransportx.WithBalancerFactory(lbx.RoundRobinFactory{}),
	)

	commandReply, err := client.Command(context.Background(), &cqrs.CommandRequest{Name: randx.HexString(10)})
	if err != nil {
		log.Fatalf("could not Command: %v", err)
	}
	log.Printf("Command: %s", commandReply)

	emptyReply, err := client.CommandEmpty(context.Background(), &cqrs.CommandRequest{Name: randx.HexString(10)})
	if err != nil {
		log.Fatalf("could not CommandEmpty: %v", err)
	}
	log.Printf("CommandEmpty: %s", emptyReply)

	queryReply, err := client.Query(context.Background(), &cqrs.QueryRequest{Name: randx.HexString(10)})
	if err != nil {
		log.Fatalf("could not Query: %v", err)
	}
	log.Printf("Query: %s", queryReply)

	queryOneOfReply, err := client.QueryOneOf(context.Background(), &cqrs.QueryRequest{Name: randx.HexString(10)})
	if err != nil {
		log.Fatalf("could not QueryOneOf: %v", err)
	}
	log.Printf("QueryOneOf: %s", queryOneOfReply)
}
