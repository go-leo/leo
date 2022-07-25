package metrics

import (
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/metric/aggregator/histogram"
	"go.opentelemetry.io/otel/sdk/metric/export"
	"go.opentelemetry.io/otel/sdk/metric/export/aggregation"
	selector "go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

type options struct {
	// AggregatorSelector 聚合器选择器, 支持在运行时为特定的度量工具选择使用的聚合器（Aggregator）类型。
	AggregatorSelector export.AggregatorSelector
	// TemporalitySelector是Exporter的子接口，用于指示处理器应该计算增量聚合还是累积聚合。
	TemporalitySelector aggregation.TemporalitySelector
	// Attributes metrics需要一些额外的信息
	Attributes []attribute.KeyValue
	// CollectPeriod 收集间隔
	CollectPeriod time.Duration
	// CollectTimeout 收集超时
	CollectTimeout time.Duration
	// PushTimeout 导出超时
	PushTimeout time.Duration
	// Prometheus 普罗米修斯系统
	Prometheus bool
	// Boundaries 直方图边界
	Boundaries []float64
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func (o *options) init() {
	if o.AggregatorSelector == nil {
		o.AggregatorSelector = selector.NewWithHistogramDistribution()
	}
	if o.TemporalitySelector == nil {
		o.TemporalitySelector = aggregation.CumulativeTemporalitySelector()
	}
}

type Option func(o *options)

// HistogramAggregatorSelector 它对histogram工具使用直方图聚合器。
func HistogramAggregatorSelector(boundaries []float64) Option {
	return func(o *options) {
		var opts []histogram.Option
		if len(boundaries) > 0 {
			// 设置配置选项的详细的边界
			opts = append(opts, histogram.WithExplicitBoundaries(boundaries))
		}
		o.AggregatorSelector = selector.NewWithHistogramDistribution(opts...)
		o.Boundaries = boundaries
	}
}

// InexpensiveAggregatorSelector 便宜的、更快，内存使用更少的聚合选择器
func InexpensiveAggregatorSelector() Option {
	return func(o *options) {
		o.AggregatorSelector = selector.NewWithInexpensiveDistribution()
	}
}

// CumulativeTemporalitySelector 累积
func CumulativeTemporalitySelector() Option {
	return func(o *options) {
		o.TemporalitySelector = aggregation.CumulativeTemporalitySelector()
	}
}

// DeltaTemporalitySelector 增量
func DeltaTemporalitySelector() Option {
	return func(o *options) {
		o.TemporalitySelector = aggregation.DeltaTemporalitySelector()
	}
}

// StatelessTemporalitySelector 无状态，不耗内存
func StatelessTemporalitySelector() Option {
	return func(o *options) {
		o.TemporalitySelector = aggregation.StatelessTemporalitySelector()
	}
}

func Attribute(attrs ...attribute.KeyValue) Option {
	return func(o *options) {
		o.Attributes = attrs
	}
}

func CollectPeriod(duration time.Duration) Option {
	return func(o *options) {
		o.CollectPeriod = duration
	}
}

func CollectTimeout(duration time.Duration) Option {
	return func(o *options) {
		o.CollectTimeout = duration
	}
}

func PushTimeout(duration time.Duration) Option {
	return func(o *options) {
		o.PushTimeout = duration
	}
}

func Prometheus() Option {
	return func(o *options) {
		o.Prometheus = true
	}
}
