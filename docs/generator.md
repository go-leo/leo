# generator

Leo提供一个 `leo` 工具 和 `protoc-gen-go-leo` protoc插件

* `leo`可以初始化项目(project)和应用(app)
* `protoc-gen-go-leo` 可以生成HTTP和gRPC的代码，还可以生成配置、状态码等代码。

# 安装

```bash
go install github.com/go-leo/leo/v3/cmd/leo@latest
go install github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo@latest
```

# leo 命令

`leo`命令有两子命令

## 1. 创建一个新项目

```
leo project -m github.com/go-leo/example
```

项目结构如下：

```
.
├── app
├── configs
├── deployments
├── docs
├── githooks
├── go.mod
├── go.sum
├── internal
├── pkg
├── scripts
├── third_party
└── tools
```

* app：应用目录，创建的应用会生成到该目录下
* configs：配置文件目录
* deployments：部署文件目录
* docs：文档目录
* githooks：git钩子目录
* go.mod,go.sum：go模块文件
* internal：私有目录, 项目内部共享但不想导出的的代码放到该目录下
* pkg：公共目录，项目内部共享且可以导出的代码放到该目录下
* scripts：脚本目录
* third_party：第三方依赖目录
* tools：工具目录

## 2. 创建一个新应用

进入项目目录

```
cd example
```

创建应用

```
leo app -n user
```

应用结构如下：

```
.
├── api
├── cmd
├── domain
├── infra
├── protoc.sh
├── service
└── ui
```

* api：api目录，存放proto定义文件
* cmd: 应用启动目录，存放main函数与子命令
* domain：领域层目录，存放领域模型代码
* infra：基础设施层目录，存放基础设施代码
* protoc.sh: protoc编译脚本
* service：服务层目录，存放业务服务代码
* ui：ui层目录，存放 user interface 代码, 代码可以通过protoc.sh生成出来

