# CQRS

CQRS splits your application (and even the database in some cases) into two different paths: **Commands** and **Queries**.

## Command side

Every operation that can trigger an side effect on the server must pass through the CQRS "command side". I like to put the `Handlers` (commands handlers and events handlers) inside the application layer because their goals are almost the
same: orchestrate domain operations (also usually using infrastructure services).

![command side](docs/images/command_side.jpg)

[//]: # (![command side]&#40;docs/images/command_side_with_events.jpg&#41;)

## Query side

Pretty straight forward, the controller receives the request, calls the related query repo and returns a DTO (defined on infrastructure layer itself).

![query side](docs/images/query_side.jpg)

# Install
```
go install github.com/go-leo/leo/v3/cmd/protoc-gen-leo-cqrs@latest
```

# Create a new project
```
├── api
│   └── demo
├── cmd
│   └── demo
│       ├── client
│       └── server
├── internal
│   └── demo
│       ├── assembler
│       ├── command
│       ├── model
│       └── query
└── third_party
    └── leo
        └── cqrs
           └── annotations.proto
```

* Copy CQRS proto [annotations.proto](leo%2Fcqrs%2Fannotations.proto) to third_party/leo/cqrs

# Write proto api
## add cqrs option at service
```protobuf
service DemoService {
  // Define command package information
  option(leo.cqrs.command) = {
    // package: the full package name of the command.
    package: "github.com/go-leo/cqrs/example/internal/demo/command"
    // relative: the package path of the command, relative to the current proto file.
    relative: "../../../internal/demo/command"
  };

  // Define command package information
  option(leo.cqrs.query) = {
    // package: the full package name of the query.
    package: "github.com/go-leo/cqrs/example/internal/demo/query"
    // relative: the package path of the command, relative to the current proto file.
    relative:  "../../../internal/demo/query"
  };
  ......
}
```
## add responsibility at method
```protobuf

service DemoService {
  ......
  // GetUsers sync get users
  rpc GetUsers (GetUsersRequest) returns (GetUsersResponse) {
    option(leo.cqrs.responsibility) = Query;
  }

  // DeleteUser sync delete user
  rpc DeleteUser (DeleteUsersRequest) returns (google.protobuf.Empty) {
    option(leo.cqrs.responsibility) = Command;
  }

  // DeleteUser async get users
  rpc AsyncGetUsers (AsyncGetUsersRequest) returns (stream AsyncGetUsersResponse) {
    option(leo.cqrs.responsibility) = Query;
  }

  // AsyncDeleteUsers async delete users
  rpc AsyncDeleteUsers(AsyncDeleteUsersRequest) returns (stream google.protobuf.Empty) {
    option(leo.cqrs.responsibility) = Command;
  }
}
```

* `option(leo.cqrs.responsibility) = Query;` means this method is a query method
* `option(leo.cqrs.responsibility) = Command;` means this method is a command method
* `returns (Response)` means this method is sync method
* `returns (stream Response)`  means this method is async method

[demo.proto](example%2Fapi%2Fdemo%2Fdemo.proto)

# generate file
```
protoc \
--proto_path=. \
--proto_path=third_party \
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
--go-grpc_opt=require_unimplemented_servers=false \
--cqrs_out=. \
--cqrs_opt=paths=source_relative \
--cqrs_opt=require_unimplemented_servers=false \
api/*/*.proto
```
will generate new files:
```
├── api
│   └── demo
│       ├── demo.cqrs.pb.go
│       ├── demo.pb.go
│       ├── demo.proto
│       └── demo_grpc.pb.go
├── cmd
│   └── demo
│       ├── client
│       └── server
├── compile.sh
├── internal
│   └── demo
│       ├── assembler
│       ├── command
│       │   ├── async_delete_users.go
│       │   ├── create_user.go
│       │   ├── delete_user.go
│       │   └── update_user.go
│       ├── model
│       └── query
│           ├── async_get_users.go
│           ├── get_user.go
│           └── get_users.go
└── third_party
    └── leo
        └── cqrs
            └── annotations.proto

```

# fill the logic
* [command](example%2Finternal%2Fdemo%2Fcommand)
* [query](example%2Finternal%2Fdemo%2Fquery)
* [model](example%2Finternal%2Fdemo%2Fmodel)
* [assembler](example%2Finternal%2Fdemo%2Fassembler)


# write the grpc server launcher
```go   
package main

import (
	"github.com/go-leo/cqrs/example/api/demo"
	"github.com/go-leo/cqrs/example/internal/demo/assembler"
	"github.com/go-leo/cqrs/example/internal/demo/command"
	"github.com/go-leo/cqrs/example/internal/demo/query"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	bus, err := demo.NewDemoServiceBus(
		command.NewCreateUser(),
		command.NewUpdateUser(),
		query.NewGetUser(),
		query.NewGetUsers(),
		command.NewDeleteUser(),
		query.NewAsyncGetUsers(),
		command.NewAsyncDeleteUsers(),
	)
	if err != nil {
		panic(err)
	}
	service := demo.NewDemoServiceCQRSService(bus, assembler.NewDemoServiceAssembler())
	demo.RegisterDemoServiceServer(s, service)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

# write grpc client call
```go
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
```

[sample](example)

