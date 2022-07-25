package otelx

import "go.opentelemetry.io/otel/sdk/metric/export/aggregation"

// TemporalitySelector是Exporter的子接口，用于指示处理器应该计算增量聚合还是累积聚合。
func NewTemporalitySelector(
	cumulative bool,
	delta bool,
	stateless bool) aggregation.TemporalitySelector {
	switch {
	case cumulative:
		// 累积聚合
		return aggregation.CumulativeTemporalitySelector()
	case delta:
		// 增量聚合
		return aggregation.DeltaTemporalitySelector()
	case stateless:
		// 无状态聚合
		return aggregation.StatelessTemporalitySelector()
	default:
		return aggregation.CumulativeTemporalitySelector()
	}
}
