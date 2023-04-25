package global

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/go-leo/otelx/metricx"
	"github.com/go-leo/otelx/tracex"
	"go.opentelemetry.io/otel/attribute"

	"github.com/go-leo/stringx"

	"github.com/go-leo/leo/common/errorx"
	"github.com/go-leo/leo/log"
	"github.com/go-leo/leo/log/zap"
)

// Config 基础配置模版
// DEPRECATED
type Config struct {
	// Server 服务Host/端口配置项
	Server Server `mapstructure:"server" json:"server" yaml:"server"`
	// Registrar 服务注册
	Registrar Registrar `mapstructure:"registrar" json:"registrar" yaml:"registrar"`
	// ServiceProviders 服务提供者配置项
	ServiceProviders ServiceProviders `mapstructure:"service_providers" json:"service_providers" yaml:"service_providers"`
	// Logger 日志配置项
	Logger LoggerConfig `mapstructure:"logger" json:"logger" yaml:"logger"`
	// Metrics 指标配置
	Metrics Metrics `mapstructure:"metrics" json:"metrics" yaml:"metrics"`
	// Trace 链路追踪配置
	Trace Trace `mapstructure:"trace" json:"trace" yaml:"trace"`
	// Management 应用管理配置
	Management Management `mapstructure:"management" json:"management" yaml:"management"`
}

// Server 服务的HTTP/gRPC信息
type (
	Server struct {
		// GRPC 服务端配置项
		GRPC GRPCServer `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
		// HTTP 服务端配置项
		HTTP HTTPServer `mapstructure:"http" json:"http" yaml:"http"`
	}
	// GRPCServer 服务端配置
	GRPCServer struct {
		// Port 端口号
		Port int `mapstructure:"port" json:"port" yaml:"port"`
		// TLS tls配置
		TLS ServerTLS `mapstructure:"tls" json:"tls" yaml:"tls"`
		// Auth
		Auth Auth `mapstructure:"auth" json:"auth" yaml:"auth"`
		// Limiter 限流器
		Limiter Limiter `mapstructure:"limiter" json:"limiter" yaml:"limiter"`
		// Recovery
		Recovery Recovery `mapstructure:"recovery" json:"recovery" yaml:"recovery"`
		// WriteBufferSize 写缓冲区大小，默认是32kB
		WriteBufferSize int `mapstructure:"write_buffer_size" json:"write_buffer_size" yaml:"write_buffer_size"`
		// ReadBufferSize 读缓冲区大小，默认是32KB
		ReadBufferSize int `mapstructure:"read_buffer_size" json:"read_buffer_size" yaml:"read_buffer_size"`
		// MaxRecvMsgSize 最大接受消息的大小
		MaxRecvMsgSize int `mapstructure:"max_recv_msg_size" json:"max_recv_msg_size" yaml:"max_recv_msg_size"`
		// MaxSendMsgSize 最大发送消息的大小
		MaxSendMsgSize int `mapstructure:"max_send_msg_size" json:"max_send_msg_size" yaml:"max_send_msg_size"`
		// MaxConcurrentStreams 最大的并发数
		MaxConcurrentStreams uint32 `mapstructure:"max_concurrent_streams" json:"max_concurrent_streams" yaml:"max_concurrent_streams"`
		// Reflection enable server reflection
		Reflection      bool            `mapstructure:"reflection" json:"reflection" yaml:"reflection"`
		KeepaliveParams KeepaliveParams `mapstructure:"keepalive_params" json:"keepalive_params" yaml:"keepalive_params"`
	}

	KeepaliveParams struct {
		MaxConnectionIdle     time.Duration `mapstructure:"max_connection_idle" json:"max_connection_idle" yaml:"max_connection_idle"`
		MaxConnectionAge      time.Duration `mapstructure:"max_connection_age" json:"max_connection_age" yaml:"max_connection_age"`
		MaxConnectionAgeGrace time.Duration `mapstructure:"max_connection_age_grace" json:"max_connection_age_grace" yaml:"max_connection_age_grace"`
		Time                  time.Duration `mapstructure:"time" json:"time" yaml:"time"`
		Timeout               time.Duration `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
		MinTime               time.Duration `mapstructure:"min_time" json:"min_time" yaml:"min_time"`
		PermitWithoutStream   bool          `mapstructure:"permit_without_stream" json:"permit_without_stream" yaml:"permit_without_stream"`
	}

	HTTPServer struct {
		// Port 端口号
		Port int `mapstructure:"port" json:"port" yaml:"port"`
		// TLS tls配置
		TLS ServerTLS `mapstructure:"tls" json:"tls" yaml:"tls"`
		// Auth
		Auth Auth `mapstructure:"auth" json:"auth" yaml:"auth"`
		// Limiter 限流器
		Limiter Limiter `mapstructure:"limiter" json:"limiter" yaml:"limiter"`
		// Recovery
		Recovery       Recovery      `mapstructure:"recovery" json:"recovery" yaml:"recovery"`
		ReadTimeout    time.Duration `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
		WriteTimeout   time.Duration `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
		IdleTimeout    time.Duration `mapstructure:"idle_timeout" json:"idle_timeout" yaml:"idle_timeout"`
		MaxHeaderBytes int           `mapstructure:"max_header_bytes" json:"max_header_bytes" yaml:"max_header_bytes"`
	}
	Auth struct {
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		// Scheme It expects the `:authorization` header to be of a certain scheme (e.g. `basic`, `bearer`), in a
		// case-insensitive format (see rfc2617, sec 1.2).
		Scheme string `mapstructure:"auth_scheme" json:"auth_scheme" yaml:"auth_scheme"`
		// Token
		Token string `mapstructure:"auth_token" json:"auth_token" yaml:"auth_token"`
	}

	Recovery struct {
		// 禁用 Recovery
		Disabled bool `mapstructure:"disabled" json:"disabled" yaml:"disabled"`
	}
	// Limiter 限流器配置，令牌桶算法
	Limiter struct {
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		// Rate 速率，限制器每秒放入rate个令牌
		Rate float64 `mapstructure:"rate" json:"rate" yaml:"rate"`
		// Bursts 限流器允许的最大突发流量，其实就是令牌桶大小
		Bursts int `mapstructure:"bursts" json:"bursts" yaml:"bursts"`
	}
)

// Registrar 注册中心
type (
	Registrar struct {
		URI string `mapstructure:"uri" json:"uri" yaml:"uri"`
	}
)

// ServiceProviders 服务提供者配置项
type (
	ServiceProviders = map[string]ServiceProvider

	// ServiceProvider 服务提供者配置信息
	ServiceProvider struct {
		// Target 服务IP地址或者是服务名
		Target string `mapstructure:"target" json:"target" yaml:"target"`
		// 服务名
		Name string `mapstructure:"name" json:"name" yaml:"name"`
		// BalancerName 负载均衡策略有：pick_first、round_robin
		BalancerName string `mapstructure:"balancer_name" json:"balancer_name" yaml:"balancer_name"`
		// Transport http or gRPC,默认gRPC
		Transport string `mapstructure:"transport" json:"transport" yaml:"transport"`
		// ClientTLS 证书
		TLS ClientTLS `mapstructure:"tls" json:"tls" yaml:"tls"`
		// CircuitBreaker 断路器
		CircuitBreaker CircuitBreaker `mapstructure:"circuit_breaker" json:"circuit_breaker" yaml:"circuit_breaker"`
	}
	// CircuitBreaker 断路器
	CircuitBreaker struct {
		// 是否启用
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		// 超时
		Timeout int `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
		// 最大并发的请求数
		MaxConcurrentRequests int `mapstructure:"max_concurrent_requests" json:"max_concurrent_requests" yaml:"max_concurrent_requests"`
		// 请求阈值  熔断器是否打开首先要满足这个条件
		RequestVolumeThreshold int `mapstructure:"request_volume_threshold" json:"request_volume_threshold" yaml:"request_volume_threshold"`
		// 熔断开启多久尝试发起一次请求
		SleepWindow int `mapstructure:"sleep_window" json:"sleep_window" yaml:"sleep_window"`
		// 误差阈值百分比
		ErrorPercentThreshold int `mapstructure:"error_percent_threshold" json:"error_percent_threshold" yaml:"error_percent_threshold"`
	}
)

// Logger 日志配置信息
type (
	LoggerConfig struct {
		// Level 日志级别 "debug" "info" "warn" "error", "panic", "fatal"
		Level string `mapstructure:"level" json:"level" yaml:"level"`
		// Fields 日志输出一些指定的键值队
		Fields  map[string]string `mapstructure:"fields" json:"fields" yaml:"fields"`
		Output  LoggerOutput      `mapstructure:"output" json:"output" yaml:"output"`
		Encoder Encoder           `mapstructure:"encoder" json:"encoder" yaml:"encoder"`
	}
	LoggerOutput struct {
		File    LoggerFile    `mapstructure:"file" json:"file" yaml:"file"`
		Console LoggerConsole `mapstructure:"console" json:"console" yaml:"console"`
	}
	LoggerConsole struct {
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	}
	LoggerFile struct {
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		// Filename 日志名称
		Filename string `mapstructure:"filename" json:"filename" yaml:"filename"`
		// Rotate 是否滚动
		Rotate bool `mapstructure:"rotate" json:"rotate" yaml:"rotate"`
		// MaxSize 是日志文件Rotate之前的最大大小(以MB为单位)。
		MaxSize int `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
		// MaxAge 是根据文件名中编码的时间戳保留旧日志文件的最大天数。默认不删除
		MaxAge int `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
		// MaxBackups 是要保留的旧日志文件的最大数量。默认保留所有，受MaxAge影响，仍然可能导致它们被删除
		MaxBackups int `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	}
	Encoder struct {
		JSON      bool `mapstructure:"json" json:"json" yaml:"json"`
		PlainText bool `mapstructure:"plain_text" json:"plain_text" yaml:"plain_text"`
	}
)

type (
	// Metrics 指标配置
	Metrics struct {
		// 是否启用指标
		Enabled    bool       `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		Controller Controller `mapstructure:"controller" json:"controller" yaml:"controller"`
		Exporter   Exporter   `mapstructure:"exporter" json:"exporter" yaml:"exporter"`
		Histogram  Histogram  `mapstructure:"histogram" json:"histogram" yaml:"histogram"`
	}
	Exporter struct {
		Prometheus Prometheus `mapstructure:"prometheus" json:"prometheus" yaml:"prometheus"`
	}
	Controller struct {
		CollectPeriod       time.Duration       `mapstructure:"collect_period" json:"collect_period" yaml:"collect_period"`
		CollectTimeout      time.Duration       `mapstructure:"collect_timeout" json:"collect_timeout" yaml:"collect_timeout"`
		PushTimeout         time.Duration       `mapstructure:"push_timeout" json:"push_timeout" yaml:"push_timeout"`
		AggregatorSelector  AggregatorSelector  `mapstructure:"aggregator_selector" json:"aggregator_selector" yaml:"aggregator_selector"`
		TemporalitySelector TemporalitySelector `mapstructure:"temporality_selector" json:"temporality_selector" yaml:"temporality_selector"`
	}
	AggregatorSelector struct {
		Inexpensive bool `mapstructure:"inexpensive" json:"inexpensive" yaml:"inexpensive"`
	}
	TemporalitySelector struct {
		Cumulative bool `mapstructure:"cumulative" json:"cumulative" yaml:"cumulative"`
		Delta      bool `mapstructure:"delta" json:"delta" yaml:"delta"`
		Stateless  bool `mapstructure:"stateless" json:"stateless" yaml:"stateless"`
	}
	Histogram struct {
		Boundaries []float64 `mapstructure:"histogram_boundaries" json:"histogram_boundaries" yaml:"histogram_boundaries"`
	}
	// Prometheus 配置
	Prometheus struct {
		// 是否启用Prometheus指标收集器
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
	}
)

// Trace 追踪系统配置
type (
	Trace struct {
		// Enabled 是否启用链路追踪
		Enabled bool `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		// Jaeger 配置
		Jaeger JaegerExporter `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
		// Zipkin 配置
		Zipkin ZipkinExporter `mapstructure:"zipkin" json:"zipkin" yaml:"zipkin"`
		// Writer 标准输入或者文件
		Writer WriterExporter `mapstructure:"writer" json:"writer" yaml:"writer"`
		// HTTP
		HTTP HTTPExporter `mapstructure:"http" json:"http" yaml:"http"`
		// GRPC
		GRPC GRPCExporter `mapstructure:"grpc" json:"grpc" yaml:"grpc"`
		// Sampler 采样策略
		Sampler Sampler `mapstructure:"sampler" json:"sampler" yaml:"sampler"`
		// ServiceName 链路追踪服务名
		ServiceName string `mapstructure:"service_name" json:"service_name" yaml:"service_name"`
	}
	GRPCExporter struct {
		Enabled            bool
		URL                string
		Insecure           bool
		Headers            map[string]string
		Timeout            time.Duration
		ReconnectionPeriod time.Duration
	}
	HTTPExporter struct {
		Enabled  bool              `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		URL      string            `mapstructure:"url" json:"url" yaml:"url"`
		Insecure bool              `mapstructure:"insecure" json:"insecure" yaml:"insecure"`
		Headers  map[string]string `mapstructure:"headers" json:"headers" yaml:"headers"`
		Timeout  time.Duration     `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
		URLPath  string            `mapstructure:"url_path" json:"url_path" yaml:"url_path"`
	}
	WriterExporter struct {
		Enabled     bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		Stdout      bool   `mapstructure:"stdout" json:"stdout" yaml:"stdout"`
		Filename    string `mapstructure:"filename" json:"filename" yaml:"filename"`
		PrettyPrint bool   `mapstructure:"pretty_print" json:"pretty_print" yaml:"pretty_print"`
	}
	JaegerExporter struct {
		Enabled  bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		URL      string `mapstructure:"url" json:"url" yaml:"url"`
		Username string `mapstructure:"username" json:"username" yaml:"username"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
	}
	ZipkinExporter struct {
		Enabled bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		URL     string `mapstructure:"url" json:"url" yaml:"url"`
	}
	Sampler struct {
		// Rate 采样概率，如果大于等于0，则全部采样，如果小于等于0，不采样
		Rate float64 `mapstructure:"rate" json:"rate" yaml:"rate"`
	}
)

type (
	ServerTLS struct {
		Enabled    bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		CertFile   string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
		KeyFile    string `mapstructure:"key_file" json:"key_file" yaml:"key_file"`
		ServerName string `mapstructure:"server_name" json:"server_name" yaml:"server_name"`
	}

	ClientTLS struct {
		Enabled    bool   `mapstructure:"enabled" json:"enabled" yaml:"enabled"`
		CertFile   string `mapstructure:"cert_file" json:"cert_file" yaml:"cert_file"`
		ServerName string `mapstructure:"server_name" json:"server_name" yaml:"server_name"`
	}
)

// Management 应用管理配置
type (
	Management struct {
		// Port 端口号
		Port int `mapstructure:"port" json:"port" yaml:"port"`
		// TLS TLS配置
		TLS ServerTLS `mapstructure:"tls" json:"tls" yaml:"tls"`
	}
)

func (c *Config) Init() {

}

func (logConf LoggerConfig) NewLogger() log.Logger {
	if stringx.IsBlank(logConf.Level) {
		logConf.Level = string(log.Info)
	}
	level := log.Level(logConf.Level)
	var opts []zap.Option
	switch {
	case logConf.Encoder.PlainText:
		opts = append(opts, zap.PlainText())
	case logConf.Encoder.JSON:
		opts = append(opts, zap.JSON())
	default:
		opts = append(opts, zap.JSON())
	}

	switch {
	case logConf.Output.File.Enabled:
		loggerFile := logConf.Output.File
		opts = append(opts, zap.File(loggerFile.Filename, loggerFile.MaxSize, loggerFile.MaxAge, loggerFile.MaxBackups))
	case logConf.Output.Console.Enabled:
		opts = append(opts, zap.Console(logConf.Output.Console.Enabled))
	default:
		opts = append(opts, zap.Console(logConf.Output.Console.Enabled))
	}

	opts = append(opts, zap.Fields(log.F{K: "hostname", V: errorx.Quiet(os.Hostname())}))
	return zap.New(zap.LevelAdapt(level), opts...)
}

func (metricConf Metrics) NewMetric(ctx context.Context) (*metricx.Metric, error) {
	// 注册prometheus官方的GoCollector和ProcessCollector
	return metricx.NewMetric(ctx)
}

func (traceConf Trace) NewTrace(ctx context.Context) (*tracex.Trace, error) {
	writerConf := traceConf.Writer
	jaegerConf := traceConf.Jaeger
	zipkinConf := traceConf.Zipkin
	httpConf := traceConf.HTTP
	gRPCConf := traceConf.GRPC
	Attributes := []attribute.KeyValue{
		attribute.Key("service.name").String(traceConf.ServiceName),
		attribute.Key("service.instance.id").String(errorx.Quiet(os.Hostname())),
	}
	switch {
	case writerConf.Enabled:
		var writer io.Writer
		switch {
		case writerConf.Stdout:
			writer = os.Stdout
		case stringx.IsNotBlank(writerConf.Filename):
			file, err := os.Create(writerConf.Filename)
			if err != nil {
				return nil, err
			}
			writer = file
		default:
			writer = io.Discard
		}
		return tracex.NewTrace(
			ctx,
			tracex.Writer(&tracex.WriterOptions{Writer: writer, PrettyPrint: writerConf.PrettyPrint}),
			tracex.SampleRate(traceConf.Sampler.Rate),
			tracex.Attributes(Attributes...),
		)
	case jaegerConf.Enabled:
		return tracex.NewTrace(
			ctx,
			tracex.Jaeger(&tracex.JaegerOptions{
				Endpoint: jaegerConf.URL,
				Username: jaegerConf.Username,
				Password: jaegerConf.Password,
			}),
			tracex.SampleRate(traceConf.Sampler.Rate),
			tracex.Attributes(Attributes...),
		)
	case zipkinConf.Enabled:
		return tracex.NewTrace(
			ctx,
			tracex.Zipkin(&tracex.ZipkinOptions{
				URL: zipkinConf.URL,
			}),
			tracex.SampleRate(traceConf.Sampler.Rate),
			tracex.Attributes(Attributes...),
		)
	case httpConf.Enabled:
		return tracex.NewTrace(
			ctx,
			tracex.HTTP(&tracex.HTTPOptions{
				Endpoint: httpConf.URL,
				Insecure: httpConf.Insecure,
				// TLSConfig:   traceConf.TL,
				Headers: httpConf.Headers,
				// Compression: traceConf.HTTP.Compression,
				// Retry:       nil,
				Timeout: httpConf.Timeout,
				URLPath: httpConf.URLPath,
			}),
			tracex.SampleRate(traceConf.Sampler.Rate),
			tracex.Attributes(Attributes...),
		)
	case gRPCConf.Enabled:
		return tracex.NewTrace(
			ctx,
			tracex.GRPC(&tracex.GRPCOptions{
				Endpoint: gRPCConf.URL,
				Insecure: gRPCConf.Insecure,
				// TLSConfig:          nil,
				Headers: gRPCConf.Headers,
				// Compressor:         "",
				// DialOptions:        nil,
				// GRPCConn:           nil,
				ReconnectionPeriod: gRPCConf.ReconnectionPeriod,
				// Retry:              nil,
				Timeout: gRPCConf.Timeout,
				// ServiceConfig:      "",
			}),
			tracex.SampleRate(traceConf.Sampler.Rate),
			tracex.Attributes(Attributes...),
		)
	}
	return nil, nil
}
