# recovery
Leo 提供了内置的panic恢复中间件，可以避免程序崩溃。

# 使用
```go
package main

import (
	"context"
	"github.com/go-leo/leo/v3"
	"github.com/go-leo/leo/v3/example/api/helloworld/v1"
	"github.com/go-leo/leo/v3/recoveryx"
	"github.com/go-leo/leo/v3/serverx/grpcserverx"
	"github.com/go-leo/leo/v3/serverx/httpserverx"
	"github.com/go-leo/leo/v3/transportx/grpctransportx"
	"github.com/go-leo/leo/v3/transportx/httptransportx"
	"github.com/gorilla/mux"
	"log"
)

func main() {
	// 添加recovery中间件
	// mdw := recoveryx.Middleware()
	mdw := recoveryx.Middleware(recoveryx.RecoveryHandler(func(ctx context.Context, p any) (err error) {
		return fmt.Errorf("panic: %v", p)
	}))

	
	// http server
	router := helloworld.AppendGreeterHttpServerRoutes(
		mux.NewRouter(),
		&server{},
		httptransportx.Middleware(mdw),
	)
	httpSrv := httpserverx.NewServer(router, httpserverx.Port(60051))

	// grpc server
	grpcSrv := grpcserverx.NewServer(grpcserverx.Port(50051))
	helloworld.RegisterGreeterServer(
		grpcSrv,
		helloworld.NewGreeterGrpcServer(
			&server{},
			grpctransportx.Middleware(mdw),
		),
	)

	if err := leo.NewApp(leo.Runner(httpSrv, grpcSrv)).Run(context.Background()); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type server struct{}

func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	panic("this is panic")
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
```


```go
// 客户端
func main() {
	grpcCli := helloworld.NewGreeterGrpcClient("localhost:50051", grpctransportx.WithDialOptions(grpc.WithTransportCredentials(insecure.NewCredentials())))
	Call(grpcCli)
	httpCli := helloworld.NewGreeterHttpClient("localhost:60051")
	Call(httpCli)
}

func Call(grpcCli helloworld.GreeterService) {
	r, err := grpcCli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "recovery"})
	if err != nil {
		st, ok := statusx.From(err)
		if ok {
			log.Printf("could not greet: %v, debugInfo: %v", err, st.DebugInfo())
		}
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
```

注意：
* 默认的recovery处理器会把panic的堆栈信息保存在`errdetails.DebugInfo`里。客户端`st.DebugInfo()`拿到。
* 建议生产环境不要用默认，自定义一个处理器，


# 代码
[recovery](../example/recovery)