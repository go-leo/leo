package leo

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gin-gonic/gin"
	"github.com/go-leo/errorx"
	"github.com/go-leo/netx/addrx"
	"github.com/go-leo/stringx"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"

	crontask "github.com/go-leo/leo/v2/cron"
	grpcserver "github.com/go-leo/leo/v2/grpc/server"
	"github.com/go-leo/leo/v2/http"
	"github.com/go-leo/leo/v2/log"
	"github.com/go-leo/leo/v2/management"
	pubsubtask "github.com/go-leo/leo/v2/pubsub"
	"github.com/go-leo/leo/v2/registry"
	"github.com/go-leo/leo/v2/runner"
)

type GRPCOptions struct {
	ServiceImpl             any
	ServiceDesc             grpc.ServiceDesc
	Port                    int
	TLSConf                 *tls.Config
	GRPCServerOptions       []grpc.ServerOption
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
	Registrar               registry.Registrar
}

type ManagementOptions struct {
	Port           int
	GinMiddlewares []gin.HandlerFunc
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	IdleTimeout    time.Duration
	MaxHeaderBytes int
	TLSConf        *tls.Config
	Routers        []management.Router
}

type CronOptions struct {
	Jobs        []*crontask.Job
	Location    *time.Location
	Middlewares []cron.JobWrapper
	Seconds     bool
	Parser      cron.ScheduleParser
}

type PubSubOptions struct {
	Jobs                 []*pubsubtask.Job
	CloseTimeout         time.Duration
	Middlewares          []message.HandlerMiddleware
	Plugins              []message.RouterPlugin
	PublisherDecorators  []message.PublisherDecorator
	SubscriberDecorators []message.SubscriberDecorator
}

type options struct {
	ID       string
	Name     string
	Version  string
	MetaData map[string]string
	Logger   log.Logger

	GRPCOpts *GRPCOptions
	HttpSrv  *http.Server

	CronOpts        *CronOptions
	PubSubOpts      *PubSubOptions
	Runnables       []runner.Runnable
	Callables       []runner.Callable
	MgmtOpts        *ManagementOptions
	ShutdownSignals []os.Signal
	ShutdownHook    func(signal os.Signal)
	RestartSignals  []os.Signal
	RestartHook     func(signal os.Signal)
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
	if o.ShutdownHook == nil {
		o.ShutdownHook = func(_ os.Signal) {}
	}
	if o.RestartHook == nil {
		o.RestartHook = func(_ os.Signal) {}
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
func GRPC(opts *GRPCOptions) Option {
	return func(o *options) {
		o.GRPCOpts = opts
	}
}

// HTTP http服务配置
func HTTP(httpSrv *http.Server) Option {
	return func(o *options) {
		o.HttpSrv = httpSrv
	}
}

// Cron 定时任务配置
func Cron(opts *CronOptions) Option {
	return func(o *options) {
		o.CronOpts = opts
	}
}

// PubSub 发布与订阅配置
func PubSub(opts *PubSubOptions) Option {
	return func(o *options) {
		o.PubSubOpts = opts
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
func Management(opts *ManagementOptions) Option {
	return func(o *options) {
		o.MgmtOpts = opts
	}
}

// ShutdownSignal app需要接受的退出信号以及回调函数
func ShutdownSignal(signals []os.Signal, hook func(signal os.Signal)) Option {
	return func(o *options) {
		o.ShutdownSignals = signals
		o.ShutdownHook = hook
	}
}

// RestartSignal app需要重启的信号以及回调函数
func RestartSignal(signals []os.Signal, hook func(signal os.Signal)) Option {
	return func(o *options) {
		o.RestartSignals = signals
		o.RestartHook = hook
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

	// 添加http服务
	if app.o.HttpSrv != nil {
		app.o.Logger.Info("add http server")
		app.executor.AddRunnable(app.o.HttpSrv)
	}

	// 等待退出
	return app.executor.Execute(ctx)
}

func (app *App) Run1(ctx context.Context) error {
	app.o.Logger.Infof("app %d starting...", os.Getpid())
	defer app.o.Logger.Infof("app %d stopping...", os.Getpid())
	ctx, app.cancel = context.WithCancel(ctx)
	// AsyncRunner管理多个Runnable和Callable运行与退出
	app.eg, ctx = errgroup.WithContext(ctx)

	// 运行自定义实现了Callable的任务
	for _, callable := range app.o.Callables {
		app.call(ctx, callable)
	}

	// 运行自定义实现了Runnable的任务
	for _, runnable := range app.o.Runnables {
		app.run(ctx, runnable)
	}
	// 启动Cron任务
	if app.o.CronOpts != nil && len(app.o.CronOpts.Jobs) > 0 {
		app.run(ctx, app.newCronTask())
	}
	// 启动PubSub任务
	if app.o.PubSubOpts != nil && len(app.o.PubSubOpts.Jobs) > 0 {
		app.run(ctx, app.newPubSubTask())
	}
	// 启动gRPC服务
	if app.o.GRPCOpts != nil {
		if err := app.startGRPCServer(ctx); err != nil {
			return err
		}
	}
	// 启动management服务
	if app.o.MgmtOpts != nil {
		if err := app.startManagementServer(ctx); err != nil {
			return err
		}
	}
	// 等待退出
	return nil
}

func (app *App) startGRPCServer(ctx context.Context) error {
	srv, err := app.newGRPCServer()
	if err != nil {
		return err
	}
	app.run(ctx, srv)
	// 注册grpc服务
	if app.o.GRPCOpts.Registrar == nil {
		return nil
	}
	serviceInfo, err := app.newServiceInfo(registry.TransportGRPC, app.o.GRPCOpts.Port)
	if err != nil {
		return err
	}
	app.run(ctx, &registrar{Registrar: app.o.GRPCOpts.Registrar, ServiceInfo: serviceInfo, Logger: app.o.Logger})
	return nil
}

func (app *App) startManagementServer(ctx context.Context) error {
	srv, err := app.newManagementServer()
	if err != nil {
		return err
	}
	app.run(ctx, srv)
	return nil
}

func (app *App) newCronTask() *crontask.Task {
	cronOpts := app.o.CronOpts
	var opts []crontask.Option
	if cronOpts.Location != nil {
		opts = append(opts, crontask.Location(cronOpts.Location))
	}
	if cronOpts.Seconds {
		opts = append(opts, crontask.Seconds())
	}
	if cronOpts.Parser != nil {
		opts = append(opts, crontask.Parser(cronOpts.Parser))
	}
	if len(cronOpts.Middlewares) > 0 {
		opts = append(opts, crontask.Middleware(cronOpts.Middlewares...))
	}
	opts = append(opts, crontask.Logger(app.o.Logger))
	return crontask.New(app.o.CronOpts.Jobs, opts...)
}

func (app *App) newPubSubTask() *pubsubtask.Task {
	subOpts := app.o.PubSubOpts
	var opts []pubsubtask.Option
	opts = append(opts, pubsubtask.Logger(app.o.Logger))
	if subOpts.CloseTimeout > 0 {
		opts = append(opts, pubsubtask.CloseTimeout(subOpts.CloseTimeout))
	}
	if len(subOpts.Middlewares) > 0 {
		opts = append(opts, pubsubtask.Middleware(subOpts.Middlewares...))
	}
	if len(subOpts.Plugins) > 0 {
		opts = append(opts, pubsubtask.Plugin(subOpts.Plugins...))
	}
	if len(subOpts.SubscriberDecorators) > 0 {
		opts = append(opts, pubsubtask.SubscriberDecorator(subOpts.SubscriberDecorators...))
	}
	if len(subOpts.PublisherDecorators) > 0 {
		opts = append(opts, pubsubtask.PublisherDecorator(subOpts.PublisherDecorators...))
	}
	return pubsubtask.New(subOpts.Jobs, opts...)
}

func (app *App) newGRPCServer() (*grpcserver.Server, error) {
	grpcOpts := app.o.GRPCOpts
	if grpcOpts.ServiceImpl == nil {
		return nil, errors.New("ServiceImpl is nil")
	}
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(grpcOpts.Port)))
	if err != nil {
		return nil, err
	}
	// 如果上面的监听的端口为0，则会随机用一个可用的端口，所以需要回填。
	grpcOpts.Port = addrx.ExtractPort(lis.Addr())
	// 组装options
	opts := []grpcserver.Option{
		grpcserver.ServerOptions(grpcOpts.GRPCServerOptions...),
		grpcserver.UnaryInterceptors(grpcOpts.UnaryServerInterceptors...),
		grpcserver.TLS(grpcOpts.TLSConf),
	}
	// 基于ServiceImpl、grpc服务的描述以及grpc的options创建 grpc server. grpc server实现了Runnable
	srv := grpcserver.New(lis, grpcserver.Service{Impl: grpcOpts.ServiceImpl, Desc: grpcOpts.ServiceDesc}, opts...)
	app.o.Logger.Infof("%s server listen at %s", srv.String(), lis.Addr())
	return srv, nil
}

func (app *App) newManagementServer() (*management.Server, error) {
	mgOpts := app.o.MgmtOpts
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(mgOpts.Port)))
	if err != nil {
		return nil, err
	}
	mgOpts.Port = addrx.ExtractPort(lis.Addr())
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}
	opts := []management.Option{
		management.GinMiddlewares(mgOpts.GinMiddlewares...),
		management.ReadTimeout(mgOpts.ReadTimeout),
		management.WriteTimeout(mgOpts.WriteTimeout),
		management.IdleTimeout(mgOpts.IdleTimeout),
		management.MaxHeaderBytes(mgOpts.MaxHeaderBytes),
		management.TLS(mgOpts.TLSConf),
		management.Routers(mgOpts.Routers...),
		management.ShutdownSignals(app.o.ShutdownSignals),
		management.RestartSignals(app.o.RestartSignals),
	}

	grpcOptions := app.o.GRPCOpts
	if grpcOptions != nil {
		target := net.JoinHostPort(host, strconv.Itoa(grpcOptions.Port))
		opts = append(opts, management.GRPCHealthCheck(target, grpcOptions.TLSConf, time.Second))
		opts = append(opts, management.GRPC(grpcOptions.ServiceDesc))
	}

	httpSrv := app.o.HttpSrv
	if httpSrv != nil {
		target := fmt.Sprintf("%s://%s%s", httpSrv.Scheme(), net.JoinHostPort(host, strconv.Itoa(httpSrv.Port())), "/")
		opts = append(opts, management.HTTPHealthCheck(target, nil, time.Second))
		opts = append(opts, management.HTTPRoutes(nil))
	}

	if app.o.CronOpts != nil {
		opts = append(opts, management.Cron(app.o.CronOpts.Jobs))
	}

	if app.o.PubSubOpts != nil {
		opts = append(opts, management.Subscriber(app.o.PubSubOpts.Jobs))
	}

	srv := management.New(lis, opts...)
	app.o.Logger.Infof("%s server listen at %s", srv.String(), lis.Addr())
	return srv, nil
}

func (app *App) newServiceInfo(transport string, port int) (*registry.ServiceInfo, error) {
	host, err := addrx.GlobalUnicastIPString()
	if err != nil {
		return nil, err
	}
	id := app.o.ID + "_" + transport + "_" + strconv.Itoa(port)
	serviceInfo := &registry.ServiceInfo{
		ID:        id,
		Name:      app.o.Name,
		Transport: transport,
		Host:      host,
		Port:      port,
		Metadata:  app.o.MetaData,
		Version:   app.o.Version,
	}
	return serviceInfo, nil
}

func (app *App) run(ctx context.Context, target runner.Runnable) {
	app.wg.Add(1)
	// 并发启动
	app.eg.Go(func() error {
		defer app.wg.Done()
		app.o.Logger.Infof("starting %s", target.String())
		return target.Start(ctx)
	})
	// 监听停止信号
	app.eg.Go(func() error {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), app.o.StopTimeout)
		defer cancel()
		err := target.Stop(ctx)
		app.o.Logger.Infof("%s is stopped", target.String())
		return err
	})
	runtime.Gosched()
}

func (app *App) call(ctx context.Context, target runner.Callable) {
	app.wg.Add(1)
	app.eg.Go(func() (err error) {
		defer app.wg.Done()
		app.o.Logger.Infof("calling %s", target.String())
		return target.Invoke(ctx)
	})
	runtime.Gosched()
}

var _ runner.Runnable = new(registrar)

type registrar struct {
	Registrar   registry.Registrar
	ServiceInfo *registry.ServiceInfo
	Logger      log.Logger
}

func (rr *registrar) String() string {
	return "registrar"
}

func (rr *registrar) Start(ctx context.Context) error {
	rr.Logger.Infof(
		"register service, name: %s, id: %s, transport: %s, address: %s",
		rr.ServiceInfo.Name, rr.ServiceInfo.ID, rr.ServiceInfo.Transport,
		net.JoinHostPort(rr.ServiceInfo.Host, strconv.Itoa(rr.ServiceInfo.Port)))
	return rr.Registrar.Register(ctx, rr.ServiceInfo)
}

func (rr *registrar) Stop(ctx context.Context) error {
	return rr.Registrar.Deregister(ctx, rr.ServiceInfo)
}
