# 元数据
go-kit是一个对通信协议无关的工具库，不同的通信方式都有自己的元数据(比如http使用Http Header 作为元数据，gRPC使用 metadata 作为元数据等)

为了解决不同的通信协议传递元数据，Leo提供了统一的元数据支持，支持跨协议传递元数据。

# 使用
## 客户端发送元数据
方式一
```go
ctx := context.Background()
ctx = metadatax.NewOutgoingContext(ctx, metadatax.Pairs("token", "1234567890"))
r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
```

方式二
```go
ctx := context.Background()
ctx = metadatax.AppendOutgoingContext(ctx, metadatax.Pairs("token", "1234567890"))
r, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "ubuntu"})
```
注意：
* `metadatax.NewOutgoingContext` 创建新的元数据到context中，会覆盖之前的元数据，不会合并context中已有的元数据
* `metadatax.AppendOutgoingContext` 追加元数据到context中, 会合并context中已有的元数据

## 服务端接受元数据
```go
md, ok := metadatax.FromIncomingContext(ctx)
if ok {
    log.Printf("token: %v", md.Get("token"))
}
```

# 代码
[metadata](../example/metadata)