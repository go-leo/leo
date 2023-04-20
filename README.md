
# Leo

## 背景

针对七猫研发团队的开发流程，提供一套基于Go语言的服务框架，脚手架工具，用于快速生成项目模板，以及项目模板中的代码生成器，代码生成器插件，代码生成器模板。

主要合并了此之前各团队研发的golaxy，leo，gin框架，降低大家开发磨合成本。

## 优势
* 支持主流框架应有功能
* 支持原有golaxy框架组件的无缝迁移
* leo框架只提供接口设计，具体实现由模块子包来实现，方便团队自定义实现

## 安装
    
```bash
go install codeup.aliyun.com/qimao/leo/leo/cmd/leo@latest
```

## 规划

* 致力于提供一套完整的开发框架，脚手架工具，代码生成器，代码生成器插件，代码生成器模板，以及相关文档
* 对于各业务团队使用依赖严格把控，不在泛滥生成新项目新分支，对于版本控制走集中化管控。

## - [leo-cli项目](https://codeup.aliyun.com/qimao/leo/leo/tree/master/cmd)

* 命令方式生成项目

## - [layout项目](https://codeup.aliyun.com/qimao/leo/layout) 

* 用于直接生成初始项目，指定mod名称
