package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/go-leo/gox/mathx/randx"
	"github.com/go-leo/leo/v3/example/api/demo"
	"google.golang.org/genproto/googleapis/api/httpbody"
	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	transports, err := demo.NewDemoGrpcClientTransports(":9090", []grpc1.DialOption{grpc1.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		panic(err)
	}
	client := demo.NewDemoGrpcClient(transports)

	createUserResp, err := client.CreateUser(context.Background(), &demo.CreateUserRequest{
		User: &demo.User{
			Name:   randx.HexString(12),
			Age:    randx.Int31n(50),
			Salary: randx.Float64(),
			Token:  randx.NumericString(16),
			Avatar: randx.WordString(16),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("CreateUser:", createUserResp)

	deleteUserResp, err := client.DeleteUser(context.Background(), &demo.DeleteUsersRequest{
		UserId: randx.Uint64(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("DeleteUser:", deleteUserResp)

	updateUserResp, err := client.UpdateUser(context.Background(), &demo.UpdateUserRequest{
		UserId: randx.Uint64(),
		User: &demo.User{
			Name:   randx.HexString(12),
			Age:    randx.Int31n(50),
			Salary: randx.Float64(),
			Token:  randx.NumericString(16),
			Avatar: randx.WordString(16),
		},
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

	b := make([]byte, 1024)
	_, _ = rand.Read(b)
	uploadUserAvatarResp, err := client.UploadUserAvatar(context.Background(), &demo.UploadUserAvatarRequest{
		UserId: randx.Uint64(),
		Avatar: &httpbody.HttpBody{
			ContentType: "image/jpb",
			Data:        b,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("UploadUserAvatar:", uploadUserAvatarResp)

	getUsersAvatarResp, err := client.GetUserAvatar(context.Background(), &demo.GetUserAvatarRequest{
		UserId: randx.Uint64(),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("GetUserAvatar:", getUsersAvatarResp)

}
