package leo

import (
	"context"
	"crypto/tls"
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
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/go-leo/leo/common/errorx"
	"github.com/go-leo/leo/common/netx"
	"github.com/go-leo/leo/common/processx"
	"github.com/go-leo/leo/common/signalx"
	"github.com/go-leo/leo/common/stringx"
	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/registry"
	"github.com/go-leo/leo/runner"
	"github.com/go-leo/leo/runner/management"
	grpcserver "github.com/go-leo/leo/runner/net/grpc/server"
	httpserver "github.com/go-leo/leo/runner/net/http/server"
	crontask "github.com/go-leo/leo/runner/task/cron"
	"github.com/go-leo/leo/runner/task/pubsub"
)

type HttpOptions struct {
	Port            int
	GRPCDialOptions []grpc.DialOption
	GRPCClient      any
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	IdleTimeout     time.Duration
	MaxHeaderBytes  int
	TLSConf         *tls.Config
	GinMiddlewares  []gin.HandlerFunc
	Routers         []httpserver.Router
}

type GRPCOptions struct {
	Port                    int
	TLSConf                 *tls.Config
	GRPCServerOptions       []grpc.ServerOption
	UnaryServerInterceptors []grpc.UnaryServerInterceptor
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
	Jobs                 []*pubsub.Job
	CloseTimeout         time.Duration
	Middlewares          []message.HandlerMiddleware
	Plugins              []message.RouterPlugin
	PublisherDecorators  []message.PublisherDecorator
	SubscriberDecorators []message.SubscriberDecorator
}

type options struct {
	ID                string
	Name              string
	Version           string
	MetaData          map[string]string
	Logger            log.Logger
	ServiceImpl       any
	GRPCDesc          *grpc.ServiceDesc
	HTTPDesc          *httpserver.ServiceDesc
	GRPCClientCreator func(cc grpc.ClientConnInterface) any
	GRPCOpts          *GRPCOptions
	HttpOpts          *HttpOptions
	Registrar         registry.Registrar
	CronOpts          *CronOptions
	PubSubOpts        *PubSubOptions
	Runnables         []runner.Runnable
	Callables         []runner.Callable
	MgmtOpts          *ManagementOptions
	ShutdownSignals   []os.Signal
	ShutdownHook      func(signal os.Signal)
	RestartSignals    []os.Signal
	RestartHook       func(signal os.Signal)
	StopTimeout       time.Duration
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

// Service 服务描述
func Service(serviceDescriber func() (any, *grpc.ServiceDesc, *httpserver.ServiceDesc, func(grpc.ClientConnInterface) any)) Option {
	return func(o *options) {
		o.ServiceImpl, o.GRPCDesc, o.HTTPDesc, o.GRPCClientCreator = serviceDescriber()
	}
}

// GRPC GRPC服务配置
func GRPC(opts *GRPCOptions) Option {
	return func(o *options) {
		o.GRPCOpts = opts
	}
}

// HTTP http服务配置
func HTTP(opts *HttpOptions) Option {
	return func(o *options) {
		o.HttpOpts = opts
	}
}

// Registrar 服务注册器
func Registrar(registrar registry.Registrar) Option {
	return func(o *options) {
		o.Registrar = registrar
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
	o           *options
	eg          *errgroup.Group
	wg          sync.WaitGroup
	stopTimeout time.Duration
	cancel      context.CancelFunc
}

func NewApp(opts ...Option) *App {
	o := &options{
		ShutdownSignals: []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT},
		RestartSignals:  []os.Signal{syscall.SIGHUP},
	}
	o.apply(opts...)
	o.init()
	return &App{o: o}
}

// Run 启动app
func (app *App) Run(ctx context.Context) error {
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
	// 运行Cron任务
	if app.o.CronOpts != nil && len(app.o.CronOpts.Jobs) > 0 {
		app.run(ctx, app.newCronTask())
	}
	// 运行PubSub任务
	if app.o.PubSubOpts != nil && len(app.o.PubSubOpts.Jobs) > 0 {
		app.run(ctx, app.newPubSubTask())
	}
	// 运行gRPC服务
	if app.o.GRPCOpts != nil {
		if err := app.startGRPCServer(ctx); err != nil {
			return err
		}
	}
	// 运行http服务
	if app.o.HttpOpts != nil {
		if err := app.startHTTPServer(ctx); err != nil {
			return err
		}
	}
	// 运行management服务
	if app.o.MgmtOpts != nil {
		if err := app.startManagementServer(ctx); err != nil {
			return err
		}
	}
	// 等待系统信号，支持关闭和重启信号
	if len(app.o.ShutdownSignals)+len(app.o.RestartSignals) > 0 {
		app.listenSignal(ctx)
	}
	// 等待退出
	return app.wait()
}

func (app *App) startGRPCServer(ctx context.Context) error {
	srv, err := app.newGRPCServer()
	if err != nil {
		return err
	}
	app.run(ctx, srv)
	// 注册grpc服务
	if app.o.Registrar == nil {
		return nil
	}
	serviceInfo, err := app.newServiceInfo(registry.TransportGRPC, app.o.GRPCOpts.Port)
	if err != nil {
		return err
	}
	app.run(ctx, &registrar{Registrar: app.o.Registrar, ServiceInfo: serviceInfo, Logger: app.o.Logger})
	return nil
}

func (app *App) startHTTPServer(ctx context.Context) error {
	srv, err := app.newHTTPServer(ctx)
	if err != nil {
		return err
	}
	app.run(ctx, srv)
	// 注册http服务
	if app.o.Registrar == nil {
		return nil
	}
	transport := registry.TransportHTTP
	if app.o.HttpOpts.TLSConf != nil {
		transport = registry.TransportHTTPS
	}
	serviceInfo, err := app.newServiceInfo(transport, app.o.HttpOpts.Port)
	if err != nil {
		return err
	}
	app.run(ctx, &registrar{Registrar: app.o.Registrar, ServiceInfo: serviceInfo, Logger: app.o.Logger})
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

func (app *App) newPubSubTask() *pubsub.Task {
	subOpts := app.o.PubSubOpts
	var opts []pubsub.Option
	opts = append(opts, pubsub.Logger(app.o.Logger))
	if subOpts.CloseTimeout > 0 {
		opts = append(opts, pubsub.CloseTimeout(subOpts.CloseTimeout))
	}
	if len(subOpts.Middlewares) > 0 {
		opts = append(opts, pubsub.Middleware(subOpts.Middlewares...))
	}
	if len(subOpts.Plugins) > 0 {
		opts = append(opts, pubsub.Plugin(subOpts.Plugins...))
	}
	if len(subOpts.SubscriberDecorators) > 0 {
		opts = append(opts, pubsub.SubscriberDecorator(subOpts.SubscriberDecorators...))
	}
	if len(subOpts.PublisherDecorators) > 0 {
		opts = append(opts, pubsub.PublisherDecorator(subOpts.PublisherDecorators...))
	}
	return pubsub.New(subOpts.Jobs, opts...)
}

func (app *App) newGRPCServer() (*grpcserver.Server, error) {
	grpcOpts := app.o.GRPCOpts
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(grpcOpts.Port)))
	if err != nil {
		return nil, err
	}
	// 如果上面的监听的端口为0，则会随机用一个可用的端口，所以需要回填。
	grpcOpts.Port = netx.ExtractPort(lis.Addr())
	// 组装options
	opts := []grpcserver.Option{
		grpcserver.ServerOptions(grpcOpts.GRPCServerOptions...),
		grpcserver.UnaryInterceptors(grpcOpts.UnaryServerInterceptors...),
		grpcserver.TLS(grpcOpts.TLSConf),
	}
	// 基于ServiceImpl、grpc服务的描述以及grpc的options创建 grpc server. grpc server实现了Runnable
	srv := grpcserver.New(lis, grpcserver.Service{Impl: app.o.ServiceImpl, Desc: app.o.GRPCDesc}, opts...)
	app.o.Logger.Infof("%s server listen at %s", srv.String(), lis.Addr())
	return srv, nil
}

func (app *App) newHTTPServer(ctx context.Context) (*httpserver.Server, error) {
	httpOpts := app.o.HttpOpts
	// 监听端口
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(httpOpts.Port)))
	if err != nil {
		return nil, err
	}
	// 如果上面的监听的端口为0，则会随机用一个可用的端口，所以需要回填。
	httpOpts.Port = netx.ExtractPort(lis.Addr())

	if app.o.GRPCOpts != nil && httpOpts.GRPCClient == nil {
		dialOptions := append([]grpc.DialOption{}, httpOpts.GRPCDialOptions...)
		if app.o.GRPCOpts.TLSConf != nil {
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(credentials.NewTLS(app.o.GRPCOpts.TLSConf)))
		} else {
			dialOptions = append(dialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}
		conn, err := grpc.DialContext(ctx, net.JoinHostPort("", strconv.Itoa(app.o.GRPCOpts.Port)), dialOptions...)
		if err != nil {
			return nil, err
		}
		httpOpts.GRPCClient = app.o.GRPCClientCreator(conn)
	}

	// 组装options
	opts := []httpserver.Option{
		httpserver.GRPCClient(httpOpts.GRPCClient),
		httpserver.ServiceDescription(app.o.HTTPDesc),
		httpserver.ReadTimeout(httpOpts.ReadTimeout),
		httpserver.WriteTimeout(httpOpts.WriteTimeout),
		httpserver.IdleTimeout(httpOpts.IdleTimeout),
		httpserver.MaxHeaderBytes(httpOpts.MaxHeaderBytes),
		httpserver.TLS(httpOpts.TLSConf),
		httpserver.Middlewares(httpOpts.GinMiddlewares...),
		httpserver.Routers(httpOpts.Routers...),
	}
	// 基于Listener和options创建 http server. http server实现了Runnable
	srv := httpserver.New(lis, opts...)
	app.o.Logger.Infof("%s server listen at %s", srv.String(), lis.Addr())
	return srv, nil
}

func (app *App) newManagementServer() (*management.Server, error) {
	mgOpts := app.o.MgmtOpts
	lis, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(mgOpts.Port)))
	if err != nil {
		return nil, err
	}
	mgOpts.Port = netx.ExtractPort(lis.Addr())
	host, err := netx.GlobalUnicastIPString()
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
		opts = append(opts, management.GRPC(app.o.GRPCDesc))
	}
	httpOptions := app.o.HttpOpts
	if httpOptions != nil {
		scheme := "http"
		if httpOptions.TLSConf != nil {
			scheme = "https"
		}
		target := fmt.Sprintf("%s://%s%s", scheme, net.JoinHostPort(host, strconv.Itoa(httpOptions.Port)), httpserver.HealthCheckPath)
		opts = append(opts, management.HTTPHealthCheck(target, httpOptions.TLSConf, time.Second))
		if app.o.HTTPDesc != nil {
			opts = append(opts, management.HTTPServiceDesc(app.o.HTTPDesc))
		}
		if len(httpOptions.Routers) > 0 {
			opts = append(opts, management.HTTPRouters(httpOptions.Routers))
		}
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
	host, err := netx.GlobalUnicastIPString()
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

func (app *App) listenSignal(ctx context.Context) {
	app.eg.Go(func() error {
		signals := append([]os.Signal{}, app.o.ShutdownSignals...)
		signals = append(signals, app.o.RestartSignals...)
		errC := make(chan error)
		go func() {
			runtime.Gosched()
			app.o.Logger.Info("app wait signals...")
			err := signalx.NewSignalWaiter(signals, 15*time.Second).
				AddHook(app.o.ShutdownHook).
				AddHook(app.o.RestartHook).
				WaitSignals().
				WaitHooksAsyncInvoked().
				WaitUntilTimeout().
				Err()
			errC <- err
			close(errC)
		}()
		select {
		case <-ctx.Done():
			return nil
		case e := <-errC:
			return e
		}
	})
}

func (app *App) wait() error {
	app.eg.Go(func() error {
		app.wg.Wait()
		app.cancel()
		return nil
	})
	err := app.eg.Wait()
	if err == nil {
		return nil
	}
	if !signalx.IsSignal(err, app.o.RestartSignals) {
		return err
	}
	if _, e := processx.StartProcess(); e != nil {
		app.o.Logger.Errorf("failed to restart process, %v", e)
		return err
	}
	app.o.Logger.Infof("restart process success")
	return err
}

var _ runner.Runnable = new(registrar)

type registrar struct {
	Registrar   registry.Registrar
	ServiceInfo *registry.ServiceInfo
	Logger      log.Logger
}

func (rr *registrar) String() string {
	return rr.Registrar.Scheme() + " registrar"
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
