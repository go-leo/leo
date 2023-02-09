package leo

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/go-leo/errorx"
	"github.com/go-leo/stringx"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"

	"github.com/go-leo/leo/v2/cron"
	"github.com/go-leo/leo/v2/grpc"
	"github.com/go-leo/leo/v2/http"
	"github.com/go-leo/leo/v2/log"
	"github.com/go-leo/leo/v2/management"
	"github.com/go-leo/leo/v2/pubsub"
	"github.com/go-leo/leo/v2/runner"
)

type options struct {
	ID       string
	Name     string
	Version  string
	MetaData map[string]string

	Logger log.Logger

	GRPCSrv *grpc.Server
	HttpSrv *http.Server

	CronTask   *cron.Task
	PubSubTask *pubsub.Task

	Runnables []runner.Runnable
	Callables []runner.Callable

	MgmtSrv         *management.Server
	ShutdownSignals []os.Signal
	RestartSignals  []os.Signal
	StopTimeout     time.Duration
}

func (o *options) init() {
	if stringx.IsBlank(o.ID) {
		switch {
		case stringx.IsNotBlank(os.Getenv("LEO_SERVICE_ID")):
			o.ID = os.Getenv("LEO_SERVICE_ID")
		default:
			o.ID = uuid.NewString()
		}
	}
	if stringx.IsBlank(o.Name) {
		switch {
		case stringx.IsNotBlank(os.Getenv("LEO_SERVICE_NAME")):
			o.Name = os.Getenv("LEO_SERVICE_NAME")
		default:
			o.Name = filepath.Base(errorx.Quiet(os.Executable()))
		}
	}
	if stringx.IsBlank(o.Version) {
		switch {
		case stringx.IsNotBlank(os.Getenv("LEO_SERVICE_VERSION")):
			o.Name = os.Getenv("LEO_SERVICE_VERSION")
		}
	}
	if o.Logger == nil {
		o.Logger = log.Discard{}
	}
	if o.StopTimeout == 0 {
		o.StopTimeout = 5 * time.Second
	}
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

// ID 服务ID
func ID(id string) Option {
	return func(o *options) {
		o.ID = id
	}
}

// Name 服务名
func Name(name string) Option {
	return func(o *options) {
		o.Name = name
	}
}

// Version 服务版本
func Version(version string) Option {
	return func(o *options) {
		o.Version = version
	}
}

// Metadata 服务其他信息
func Metadata(metaData map[string]string) Option {
	return func(o *options) {
		o.MetaData = metaData
	}
}

// Logger 日志组件
func Logger(logger log.Logger) Option {
	return func(o *options) {
		o.Logger = logger
	}
}

// GRPC GRPC服务配置
func GRPC(grpcSrv *grpc.Server) Option {
	return func(o *options) {
		o.GRPCSrv = grpcSrv
	}
}

// HTTP http服务配置
func HTTP(httpSrv *http.Server) Option {
	return func(o *options) {
		o.HttpSrv = httpSrv
	}
}

// Cron 定时任务
func Cron(task *cron.Task) Option {
	return func(o *options) {
		o.CronTask = task
	}
}

// PubSub 发布与订阅任务
func PubSub(task *pubsub.Task) Option {
	return func(o *options) {
		o.PubSubTask = task
	}
}

// Runnable 其他实现了Runnable接口的程序
func Runnable(r ...runner.Runnable) Option {
	return func(o *options) {
		o.Runnables = append(o.Runnables, r...)
	}
}

// Callable 其他实现了Callable接口的程序
func Callable(c ...runner.Callable) Option {
	return func(o *options) {
		o.Callables = append(o.Callables, c...)
	}
}

// Management 有助于对应用程序进行监控和管理，通过restful api请求来监管、审计、收集应用的运行情况
func Management(srv *management.Server) Option {
	return func(o *options) {
		o.MgmtSrv = srv
	}
}

// ShutdownSignal 关闭信号
func ShutdownSignal(signals []os.Signal) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
	}
}

// RestartSignal 重启信号
func RestartSignal(signals []os.Signal) Option {
	return func(o *options) {
		o.RestartSignals = signals
	}
}

func StopTimeout(timeout time.Duration) Option {
	return func(o *options) {
		o.StopTimeout = timeout
	}
}

type App struct {
	o        *options
	eg       *errgroup.Group
	wg       sync.WaitGroup
	cancel   context.CancelFunc
	executor *runner.Executor
}

func NewApp(opts ...Option) *App {
	o := &options{
		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
		RestartSignals:  []os.Signal{syscall.SIGHUP},
	}
	o.apply(opts...)
	o.init()
	executor := runner.NewExecutor(
		runner.Logger(o.Logger),
		runner.ShutdownSignals(o.ShutdownSignals),
		runner.RestartSignals(o.RestartSignals),
		runner.StopTimeout(o.StopTimeout),
	)
	return &App{o: o, executor: executor}
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
	app.o.Logger.Infof("app %d starting...", os.Getpid())
	defer app.o.Logger.Infof("app %d stopping...", os.Getpid())

	// 添加自定义实现了Callable的任务
	for _, callable := range app.o.Callables {
		app.o.Logger.Info("add callable")
		app.executor.AddCallable(callable)
	}

	// 添加自定义实现了Runnable的任务
	for _, runnable := range app.o.Runnables {
		app.o.Logger.Info("add runnable")
		app.executor.AddRunnable(runnable)
	}

	// 添加http服务
	if app.o.HttpSrv != nil {
		app.o.Logger.Info("add http server")
		app.o.HttpSrv.SetServiceInfo(app)
		app.executor.AddRunnable(app.o.HttpSrv)
	}

	// 添加gRPC服务
	if app.o.GRPCSrv != nil {
		app.o.Logger.Info("add grpc server")
		app.o.GRPCSrv.SetServiceInfo(app)
		app.executor.AddRunnable(app.o.GRPCSrv)
	}

	// 添加Cron任务
	if app.o.CronTask != nil {
		app.o.Logger.Info("add cron task")
		app.executor.AddRunnable(app.o.CronTask)
	}

	// 添加PubSub任务
	if app.o.PubSubTask != nil {
		app.o.Logger.Info("add pubsub task")
		app.executor.AddRunnable(app.o.PubSubTask)
	}

	// 添加management服务
	if app.o.MgmtSrv != nil {
		app.o.Logger.Info("add management server")
		app.executor.AddRunnable(app.o.MgmtSrv)
	}

	// 等待退出
	return app.executor.Execute(ctx)
}

func (app *App) ID() string {
	return app.o.ID
}

func (app *App) Name() string {
	return app.o.Name
}

func (app *App) Version() string {
	return app.o.Version
}

func (app *App) MetaData() map[string]string {
	return app.o.MetaData
}
