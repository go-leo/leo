package main

import (
	"context"

	"codeup.aliyun.com/qimao/leo/leo"
	"codeup.aliyun.com/qimao/leo/leo/example/runner"
)

func main() {
	app := leo.NewApp(
		leo.Runner(runner.PrintHello{}),
		leo.Runner(runner.LoopPrintHello{}),
	)
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
