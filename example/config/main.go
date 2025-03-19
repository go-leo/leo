package main

import (
	"context"
	"fmt"
	"github.com/go-leo/gox/contextx"
	"github.com/go-leo/leo/v3/configx"
	"github.com/go-leo/leo/v3/configx/environx"
	"github.com/go-leo/leo/v3/configx/filex"
	"github.com/go-leo/leo/v3/example/api/configs/v1"
	"os"
	"time"
)

var content = `
{
  "grpc": {
    "addr": "localhost",
    "port": 10
  },
  "redis": {
    "network": "tcp",
    "addr": "localhost",
    "password": "%d",
    "db": 20
  }
}
`

var filename = "/tmp/config.json"

func init() {
	// Mock文件配置
	if err := os.WriteFile(filename, []byte(fmt.Sprintf(content, time.Now().Unix())), 0644); err != nil {
		panic(err)
	}
	// Mock环境变量配置
	os.Setenv("LEO_RUN_ENV", "this is leo run env")
}

func main() {
	// 文件资源
	fileRes := &filex.Resource{
		Filename: filename,
	}
	// 环境变量资源
	envRes := &environx.Resource{
		Prefix: "LEO",
	}

	// 加载配置
	ctx := context.Background()
	if err := configs.LoadApplicationConfig(
		ctx,
		configx.WithResource(fileRes, envRes),
	); err != nil {
		panic(err)
	}
	// 获取配置
	fmt.Println(configs.GetApplicationConfig())

	// 监听配置, stop停止监听
	ctx, stop := context.WithCancel(ctx)
	defer stop()
	if err := configs.WatchApplicationConfig(ctx, configx.WithResource(fileRes, envRes)); err != nil {
		panic(err)
	}

	// 加载并监听配置
	if err := configs.LoadAndWatchApplicationConfig(ctx, configx.WithResource(fileRes, envRes)); err != nil {
		panic(err)
	}

	go func() {
		// 模拟配置文件变化和环境变量变化
		for {
			if time.Now().Second()%2 == 0 {
				file, err := os.OpenFile(filename, os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				_, err = file.WriteString(fmt.Sprintf(content, time.Now().Unix()))
				if err != nil {
					panic(err)
				}
			} else {
				os.Setenv("LEO_RUN_ENV", fmt.Sprintf("this is leo run env, %s", time.Now().String()))
			}
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			// 获取配置
			fmt.Println(configs.GetApplicationConfig())
		}
	}()

	ctx, cancelFunc := contextx.Signal(os.Interrupt)
	defer cancelFunc()
	<-ctx.Done()
}
