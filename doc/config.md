# 配置
一个应用程序可能从多个源获取配置。例如，应用程序可能从环境变量获取配置，从文件获取配置，从配置服务(例如: consul、nacos等)获取配置等。

Leo的configx包就是帮助开发者，从多个媒介中加载配置和监听配置。

# 配置源
Leo当前内置了四种源开箱即用：
1. [环境变量](https://github.com/go-leo/leo/tree/feature/v3/configx/environx)
2. [文件](https://github.com/go-leo/leo/tree/feature/v3/configx/filex)
3. [consul](https://github.com/go-leo/leo/tree/feature/v3/configx/consulx)
4. [nacos](https://github.com/go-leo/leo/tree/feature/v3/configx/nacosx)

# 配置的格式
Leo当前支持了四种常用的配置格式:
1. [Env](https://github.com/go-leo/leo/tree/feature/v3/configx/envx)
1. [JSON](https://github.com/go-leo/leo/tree/feature/v3/configx/jsonx)
2. 