# Basic Auth
Basic Auth 是一种简单的身份验证机制，它使用用户名和密码来验证用户的身份。

Leo提供了Basic认证中间件[basicx](../authx/basicx)

更多关于 basic auth 的文档请参考
[The 'Basic' HTTP Authentication Scheme](https://datatracker.ietf.org/doc/html/rfc7617)

# Usage
## gRPC
### 服务端
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// basic 中间件
	mdw := basicx.Server("ubuntu", "mint")
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
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// success
	// basic 中间件
	mdw := basicx.Client("ubuntu", "mint")
	client := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// basic 中间件
	mdw = basicx.Client("ubuntu", "redhat")
	client = helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithMiddleware(mdw),
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	r, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}
```

## Http
### 服务端
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/gorilla/mux"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	// basic 中间件
	mdw := basicx.Server("ubuntu", "mint")
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
	"github.com/go-leo/leo/v3/authx/basicx"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"log"
)

func main() {
	// success
	// basic 中间件
	mdw := basicx.Client("ubuntu", "mint")
	client := helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	r, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "ubuntu"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

	// error
	// basic 中间件
	mdw = basicx.Client("ubuntu", "redhat")
	client = helloworld.NewGreeterHttpClient("localhost:60051", httptransportx.WithMiddleware(mdw))
	r, err = client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "mint"})
	if err == nil {
		panic(err)
	}
	fmt.Println(err)
}

```

# 代码
[basic](../example/basic)