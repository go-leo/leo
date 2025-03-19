package main

import (
	"flag"
	"fmt"
	"github.com/go-leo/leo/v3/cmd"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/config"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/core"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/grpc"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/http"
	"github.com/go-leo/leo/v3/cmd/protoc-gen-go-leo/gen/status"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	showVersion := flag.Bool("version", false, "print the version and exit")
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-leo %v\n", cmd.Version)
		return
	}

	var flags flag.FlagSet
	options := &protogen.Options{ParamFunc: flags.Set}
	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate(plugin)
	})
}

func generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		// 错误状态码生成
		statusGenerator := status.NewGenerator(plugin, file)
		statusGenerator.Generate()

		// 配置生成
		configGenerator := config.NewGenerator(plugin, file)
		configGenerator.Generate()

		// 核心代码生成
		coreGenerator, err := core.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := coreGenerator.Generate(); err != nil {
			return err
		}

		// grpc代码生成
		grpcGenerator, err := grpc.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := grpcGenerator.Generate(); err != nil {
			return err
		}

		// http代码生成
		httpGenerator, err := http.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := httpGenerator.Generate(); err != nil {
			return err
		}
	}
	return nil
}
