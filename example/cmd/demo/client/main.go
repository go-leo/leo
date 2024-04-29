package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-leo/cqrs/example/api/demo"
	"github.com/go-leo/gox/mathx/randx"
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
	client := demo.NewDemoServiceClient(conn)

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
		Name:   "tom",
		Age:    30,
		Salary: 30000,
		Token:  "4108475619",
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
		Name: "jax",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("DeleteUser:", deleteUserResp)

	asyncGetUsersRespStream, err := client.AsyncGetUsers(context.Background(), &demo.AsyncGetUsersRequest{
		PageNo:   1,
		PageSize: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("AsyncGetUsers wait...")
	asyncGetUsersResp, err := asyncGetUsersRespStream.Recv()
	if err != nil {
		panic(err)
	}
	fmt.Println("AsyncGetUsers:", asyncGetUsersResp)

	asyncDeleteUsersStream, err := client.AsyncDeleteUsers(context.Background(), &demo.AsyncDeleteUsersRequest{
		Names: []string{"jax", "tom", "jerry"},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("AsyncDeleteUsers wait...")
	asyncDeleteUsersResp, err := asyncDeleteUsersStream.Recv()
	if err != nil {
		panic(err)
	}
	fmt.Println("AsyncDeleteUsers:", asyncDeleteUsersResp)

}
