
## About Leo Cli

* 命令行脚手架
* 代码生成器
* 代码生成器插件
* 代码生成器模板

### 安装leo-cli

```shell
go install codeup.aliyun.com/qimao/leo/leo/cmd/leo@latest
```

### 查看leo-cli帮助
* leo new -h


## 安装使用流程
```bash
# 1. 用脚手架生成项目模板
leo new codeup.aliyun.com/qimao/XXXX(期望的mod path名称) -p qimao/XXXX(期望的目录位置)
# 2. 项目模板中的依赖项目
cd codeup.aliyun.com/qimao/XXXX && make init
# 3. 生成代码
make api
# 4. 运行代码
make run

```

会自动在当前目录下生成一个项目模板，包并包含mod路径，打开项目后，可以直接运行，请自行添加git仓库信息。

```bash

git remote add origin XXXX
git add . && git commit -m "init" && git push -u origin master

```

## 模板项目运行说明

注：该项目需要本地安装go1.20及以上版本，依赖于go版本源码中slog，和 sync/atomic新特性




  

 