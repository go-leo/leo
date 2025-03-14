package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	kitloglevel "github.com/go-kit/log/level"
	"github.com/go-leo/leo/v3/logx"
	"os"
)

func main() {
	// new logger
	file, err := os.OpenFile("/tmp/example.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logger := logx.New(
		file,
		logx.JSON(),
		logx.Level(kitloglevel.DebugValue()),
		logx.Timestamp(),
		logx.Caller(0),
		logx.Sync(),
	)
	logger.Log("msg", "this logx.New() message")

	// inject key value pair to context
	ctx := logx.KeyValsExtractorInjector(context.Background(), "trace_id", "123456", "parent_id", "abcdefg")
	ctx = logx.KeyValsExtractorInjector(ctx, "span_id", "987654")
	logx.Infoln(ctx, "this is print extra key value pairs")

	fmt.Println("global debug level")
	logx.Debug(context.Background(), "method", "Debug")
	logx.Debugf(context.Background(), "method: %s", "Debugf")
	logx.Debugln(context.Background(), "method", "Debugln")

	logx.Info(context.Background(), "method", "Info")
	logx.Infof(context.Background(), "method: %s", "Infof")
	logx.Infoln(context.Background(), "method", "Infoln")

	logx.Warn(context.Background(), "method", "Warn")
	logx.Warnf(context.Background(), "method: %s", "Warnf")
	logx.Warnln(context.Background(), "method", "Warnln")

	logx.Error(context.Background(), "method", "Error")
	logx.Errorf(context.Background(), "method: %s", "Errorf")
	logx.Errorln(context.Background(), "method", "Errorln")

	fmt.Println("replace global to info level")
	logx.Replace(logx.New(os.Stdout, logx.Logfmt(), logx.Level(level.InfoValue()), logx.Timestamp(), logx.Caller(2), logx.Sync()))
	logx.Debug(context.Background(), "method", "Debug")
	logx.Debugf(context.Background(), "method: %s", "Debugf")
	logx.Debugln(context.Background(), "method", "Debugln")

	logx.Info(context.Background(), "method", "Info")
	logx.Infof(context.Background(), "method: %s", "Infof")
	logx.Infoln(context.Background(), "method", "Infoln")

	logx.Warn(context.Background(), "method", "Warn")
	logx.Warnf(context.Background(), "method: %s", "Warnf")
	logx.Warnln(context.Background(), "method", "Warnln")

	logx.Error(context.Background(), "method", "Error")
	logx.Errorf(context.Background(), "method: %s", "Errorf")
	logx.Errorln(context.Background(), "method", "Errorln")

	fmt.Println("replace global to warn level")
	logx.Replace(logx.New(os.Stdout, logx.Logfmt(), logx.Level(level.WarnValue()), logx.Timestamp(), logx.Caller(2), logx.Sync()))
	logx.Debug(context.Background(), "method", "Debug")
	logx.Debugf(context.Background(), "method: %s", "Debugf")
	logx.Debugln(context.Background(), "method", "Debugln")

	logx.Info(context.Background(), "method", "Info")
	logx.Infof(context.Background(), "method: %s", "Infof")
	logx.Infoln(context.Background(), "method", "Infoln")

	logx.Warn(context.Background(), "method", "Warn")
	logx.Warnf(context.Background(), "method: %s", "Warnf")
	logx.Warnln(context.Background(), "method", "Warnln")

	logx.Error(context.Background(), "method", "Error")
	logx.Errorf(context.Background(), "method: %s", "Errorf")
	logx.Errorln(context.Background(), "method", "Errorln")

	fmt.Println("replace global to error level")
	logx.Replace(logx.New(os.Stdout, logx.Logfmt(), logx.Level(level.ErrorValue()), logx.Timestamp(), logx.Caller(2), logx.Sync()))
	logx.Debug(context.Background(), "method", "Debug")
	logx.Debugf(context.Background(), "method: %s", "Debugf")
	logx.Debugln(context.Background(), "method", "Debugln")

	logx.Info(context.Background(), "method", "Info")
	logx.Infof(context.Background(), "method: %s", "Infof")
	logx.Infoln(context.Background(), "method", "Infoln")

	logx.Warn(context.Background(), "method", "Warn")
	logx.Warnf(context.Background(), "method: %s", "Warnf")
	logx.Warnln(context.Background(), "method", "Warnln")

	logx.Error(context.Background(), "method", "Error")
	logx.Errorf(context.Background(), "method: %s", "Errorf")
	logx.Errorln(context.Background(), "method", "Errorln")
}
