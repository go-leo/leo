# retry
Leo 提供了重试中间件，客户端请求在失败时，会进行重试。
* 只有当服务端返回含有 `errdetails.RetryInfo`信息的错误时才会重试.
* 如果RetryDelay为负数，则全链路禁止重试。

# 使用
```go
func main() {
	httpCli := helloworld.NewGreeterHttpClient(
		"localhost:60051",
		httptransportx.WithMiddleware(
			retryx.Middleware(retryx.MaxAttempts(3), retryx.Backoff(retryx.Exponential())),
		),
	)
	Call(httpCli)
}

func Call(grpcCli helloworld.GreeterService) {
	r, err := grpcCli.SayHello(context.Background(), &helloworld.HelloRequest{Name: "retry"})
	if err != nil {
		st, ok := statusx.From(err)
		if ok {
			log.Printf("could not greet: %v, retryInfo: %v", err, st.RetryInfo())
		}
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
```
# 代码
[retry](../example/retry)