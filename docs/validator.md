# Validator 参数校验

PGV（protoc-gen-validate）是一个 protoc 插件，用于生成多语言的消息验证器。虽然 Protocol Buffers 能够保证结构化数据的类型，但它们无法强制执行值的语义规则。PGV 插件通过向 protoc 生成的代码中添加支持，来验证这些约束。

Leo提供中间件，来支持PGV参数校验。

# 安装
```
go install  github.com/envoyproxy/protoc-gen-validate/cmd/protoc-gen-validate-go
```

# proto定义校验规则
```go
message HelloRequest {
  string name = 1 [(validate.rules).string.min_len = 1];
}
```
更多检验规则见[constraint-rules](https://github.com/bufbuild/protoc-gen-validate?tab=readme-ov-file#constraint-rules)

# 使用validator中间件
服务端校验：
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/validatorx"
	"log"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	// 添加 validator 中间件
	mdw := validatorx.Middleware()
	greeterGrpcServer := helloworld.NewGreeterGrpcServer(&server{}, grpctransportx.Middleware(mdw))
	helloworld.RegisterGreeterServer(grpcSrv, greeterGrpcServer)
	if err := leo.NewApp(leo.Runner(grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

客户端校验：
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/validatorx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	// 添加 validator 中间件
	mdw := validatorx.Middleware()
	c := helloworld.NewGreeterGrpcClient(
		"localhost:50051",
		grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())),
		grpctransportx.WithMiddleware(mdw),
	)
	Call(c, "")
	Call(c, "leo")
}

func Call(c helloworld.GreeterService, name string) {
	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Printf("could not greet: %v", err)
	} else {
		log.Printf("Greeting: %s", r.GetMessage())
	}
}
```

# 代码
[validator](../example/validator)