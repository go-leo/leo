package main

import (
	"context"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/cqrs/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	client := cqrs.NewCqrsGrpcClient(
		"localhost:50051",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
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
