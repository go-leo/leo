package main

import (
	"codeup.aliyun.com/qimao/leo/leo"
	"codeup.aliyun.com/qimao/leo/leo/example/runner"
	"codeup.aliyun.com/qimao/leo/leo/log"
	"codeup.aliyun.com/qimao/leo/leo/log/slog"
	"context"
)

func main() {
	app := leo.NewApp(
		leo.Logger(slog.New(slog.LevelAdapt(log.LevelDebug))),
		leo.Runner(runner.PrintHello{}),
		leo.Runner(runner.LoopPrintHello{}),
	)
	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
