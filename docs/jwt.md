# JWT Auth
一种基于JSON的轻量级的身份验证和授权机制，用于在客户端和服务器之间安全地传输信息。

Leo提供了JWT认证中间件[jwtx](../authx/jwtx)

更多关于 jwt 的文档请参考
[https://jwt.io/introduction](https://jwt.io/introduction)

# 使用
## gRPC 示例
### 服务端
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 从上下文中获取jwt信息
	claims, _ := jwtx.ClaimsFromContext(ctx)
	fmt.Println(claims)
	// 从上下文中获取jwt token
	token, _ := jwtx.TokenFromContext(ctx)
	fmt.Println(token)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// jwt 中间件
	mdw := jwtx.Server(func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil })
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	service := helloworld.NewGreeterGrpcServer(&server{}, grpctransportx.Middleware(mdw))
	helloworld.RegisterGreeterServer(grpcSrv, service)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

```

### 客户端
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// success
	// jwt 中间件
	mdw := jwtx.Client([]byte("jwt_key_secret"))
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	// 向ctx中注入jwt信息
	ctx := jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// jwt 中间件
	mdw = jwtx.Client([]byte("wrong_jwt_key_secret"))
	client = helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	// 向ctx中注入jwt信息
	ctx = jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
```

## Http 示例
### 服务端
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 从上下文中获取jwt信息
	claims, _ := jwtx.ClaimsFromContext(ctx)
	fmt.Println(claims)
	// 从上下文中获取jwt token
	token, _ := jwtx.TokenFromContext(ctx)
	fmt.Println(token)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// jwt 中间件
	mdw := jwtx.Server(func(token *jwt.Token) (interface{}, error) { return []byte("jwt_key_secret"), nil })
	router := mux.NewRouter()
	router = helloworld.AppendGreeterHttpServerRoutes(router, &server{}, httptransportx.Middleware(mdw))
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))
	if err := leo.NewApp(leo.Runner(httpSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```
### 客户端
```go
package main

import (
	"context"
	"fmt"
	"github.com/go-leo/leo/v3/authx/jwtx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

func main() {
	// success
	// jwt 中间件
	mdw := jwtx.Client([]byte("jwt_key_secret"))
	client := helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	// 向ctx中注入jwt信息
	ctx := jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// jwt 中间件
	mdw = jwtx.Client([]byte("wrong_jwt_key_secret"))
	client = helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	// 向ctx中注入jwt信息
	ctx = jwtx.NewContentWithClaims(context.Background(), jwt.MapClaims{"user_id": "123456"})
	r, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
```

# 代码
[jwt](../example/jwt)