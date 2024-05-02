package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/demo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := demo.NewDemoServiceHTTPClient(conn, []endpoint.Middleware{})

	createUserResp, err := client.CreateUser(context.Background(), &demo.CreateUserRequest{
		Name:   randx.HexString(12),
		Age:    randx.Int31n(50),
		Salary: float64(randx.Int31n(100000)),
		Token:  randx.NumericString(16),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateUser:", createUserResp)

	updateUserResp, err := client.UpdateUser(context.Background(), &demo.UpdateUserRequest{
		Name:   randx.HexString(12),
		Age:    randx.Int31n(50),
		Salary: float64(randx.Int31n(100000)),
		Token:  randx.NumericString(16),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("UpdateUser:", updateUserResp)

	getUserResp, err := client.GetUser(context.Background(), &demo.GetUserRequest{
		UserId: randx.Uint64(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("GetUser:", getUserResp)

	getUsersResp, err := client.GetUsers(context.Background(), &demo.GetUsersRequest{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("GetUsers:", getUsersResp)

	deleteUserResp, err := client.DeleteUser(context.Background(), &demo.DeleteUsersRequest{
		UserId: randx.Uint64(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("DeleteUser:", deleteUserResp)

}
