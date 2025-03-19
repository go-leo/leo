# Leo
Leo 是一个基于 [go-kit](https://github.com/go-kit/kit) 的微服务工具，简化了基于go-kit开发的繁琐的工作。

Leo 提供一些列 proto 插件，可以生成基于 go-kit 的 HTTP 和 gRPC 的代码。

# Leo 的优点
* 模块化：基于 Go-kit, 设计时考虑了模块化，允许开发人员根据具体的使用情况选择所需的组件。
* 传输协议无关：它支持多种传输协议（HTTP、gRPC 等），使其在不同的通信需求中具有灵活性。
* 服务发现：Leo 和 Go-kit 提供了内置的服务发现支持，这对于微服务架构至关重要。
* 负载均衡：包含负载均衡机制，以便在多个服务实例之间分配请求。
* 框架本身和业务代码保持一种低耦合的状态
* 中间件支持：一套通用的`middleware`，使之与`HTTP`和`gRPC`等传输协议无关
* 仪表化：它与监控和日志记录工具集成良好，提供对服务性能和健康状况的可见性。
* 标准化：推广最佳实践和标准化，使得维护和扩展微服务变得更容易。

# 功能组件
* [code generator](docs/generator.md)
  * 生成gRPC、Http、config、status代码。
  * 生成一套符合微服务和DDD思想的代码结构。
* [服务发现](docs/sd.md)
  * 扩展go-kit的服务发现功能，支持多种注册中心(consul、nacos)
* [流量染色](docs/stain.md)
  * 支持流量染色
* [限流](docs/ratelimit.md)
  * SlideWindow 滑动窗口限流
  * LeakyBucket 漏桶限流
  * TokenBucket 令牌桶限流
  * Redis Redis分布式限流
  * BBR 基于BBR的限流
* [熔断](docs/circuitbreaker.md)
  * google sre 熔断算法
  * hystrix 熔断器
  * sony go breaker
* [负载均衡](docs/loadbalance.md)
  * 扩展go-kit的负载均衡功能，支持多种负载均衡算法(随机、轮询、一致性哈希)
* [超时](docs/timeout.md)
  * 除了gRPC天然支持超时，HTTP也支持同样支持
* [重试](docs/retry.md)
  * 支持客户端失败重试。
* [配置](docs/config.md)
  * 支持从多种配置源(consul、nacos、环境变量、文件)获取配置
  * 支持监听配置变化，支持配置热加载
  * protobuf 定义配置格式，严格控制配置格式
* [状态](docs/status.md)
  * 基于 googleapi 错误规范实现，使用简单的协议无关错误模型，这使我们能够在不同的API，API协议（如gRPC或HTTP）以及错误上下文（例如，异步，批处理或工作流错误）中获得一致的体验。
* [元数据](docs/metadata.md)
  * Leo提供了一个元数据支持，支持跨通信方式传递元数据。
* [健康检查](docs/health.md)
  * gRPC和HTTP都支持健康检查
  * 支持自定义其他系统（比如redis、mysql等）的健康检查
* [日志](docs/log.md)
  * go-kit 的日志功能
* [监控](docs/opentelemetry.md)
  * 使用 OpenTelemetry 提供的监控方案
* [链路追踪](docs/opentelemetry.md)
  * 使用 OpenTelemetry 提供的链路追踪方案
* [参数校验](docs/validator.md)
  * 支持请求参数的自动校验(github.com/envoyproxy/protoc-gen-validate)
  * 避免手动检查代码
  * 支持自定义校验器
* [Panic恢复](docs/recovery.md)
  * 避免程序崩溃
* [JWT Auth](docs/jwt.md)
* [Basic Auth](docs/basic.md)
* [中间件](docs/middleware.md)
  * 除了内置限流、校验、日志、监控等中间件，go-kit的所有中间件都支持
