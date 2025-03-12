# CQRS
CQRS是一个架构模式，它将应用程序分解为两个不同的路径：命令和查询。
## 命令侧
每个可以触发服务器上的副作用的操作都必须通过CQRS“命令侧”。
![command side](images/command_side.jpg)

## 查询侧
每个可以触发服务器上的查询操作都必须通过CQRS“查询侧”。
![query side](images/query_side.jpg)

# Leo CQRS 定义
1. 通过`protobuf`定义服务
2. 一个 `rpc method` 是命令还是查询，可以通过 `Output Message` 决定
    * 如果 `Output Message`没有一个参数，则是认为是命令
    * 如果 `Output Message`有至少一个参数(普通参数或者oneof参数)，则是认为是查询。

# 案例
```go

```

## CQRS

CQRS splits your application (and even the database in some cases) into two different paths: **Commands** and **Queries**.

### Command side

Every operation that can trigger an side effect on the server must pass through the CQRS "command side". I like to put the `Handlers` (commands handlers and events handlers) inside the application layer because their goals are almost the same: orchestrate domain operations (also usually using infrastructure services).

![command side](images/command_side.jpg)

### Query side

Pretty straight forward, the controller receives the request, calls the related query repo and returns a DTO (defined on infrastructure layer itself).

![query side](images/query_side.jpg)

