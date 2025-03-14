# 日志
Leo 使用 go-kit 的日志功能，支持多种日志格式，支持多种日志输出方式，支持多种日志级别。

# 使用
使用日志最方便的方式是使用全局日志，可以直接使用`logx`包下的方法。

## 日志等级
go-kit提供4中的日志级别，debug、info、warn、error，info是默认的日志级别

## 输出


## 全局日志
Leo 提供了全局日志。
```go
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
```
注意：
* logx.XXX()打印key-value键值对形式的日志
* logx.XXXf()打印字符串格式化的日志，底层使用`fmt.Sprintf()`
* logx.XXXln()打印字符串格式化的日志，底层使用`fmt.Sprint()`
* 第一个参数的`context.Context`，用于传递日志上下文，可以传递一些额外的信息，比如请求ID、用户ID等,用法如下。

## 全局日志输出Context里的信息
```go
// 创建一个上下文，并注入一些额外的信息
ctx := logx.KeyValsExtractorInjector(context.Background(), "trace_id", "123456", "parent_id", "abcdefg")
ctx = logx.KeyValsExtractorInjector(ctx, "span_id", "987654")
logx.Infoln(ctx, "this is print extra key value pairs")
```
注意：
* logx.KeyValsExtractorInjector()方法，用于将额外的信息注入到上下文中，以便在日志中打印出来。
* logx.KeyValsExtractorInjector()方法可以调用多次，也可以一次传多个键值对。

## 修改全局日志
```go
logx.Replace(logx.New(os.Stdout, logx.Logfmt(), logx.Level(level.InfoValue()), logx.Timestamp(), logx.Caller(2), logx.Sync()))
```
注意:
* logx.Replace可以替换掉全局日志
* logx.New()可以创建一个日志实例。

## 自定义日志
除了使用全局日志外，还可以创建一个日志实例。
### 创建
```go
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
```
注意:
* Logger只有一个方法`Log(keyvals ...interface{}) error`,输出key-value键值对形式的日志

# 代码
[log](../example/log)