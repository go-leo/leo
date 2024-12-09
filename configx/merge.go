package configx

import (
	"github.com/go-leo/gox/errorx"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ Merger = (*merger)(nil)

// merger Merger 接口默认实现，用于合并多个structpb.Struct对象。
// 如果两个struct对象有相同的key，则使用后面的值覆盖前面的值。
type merger struct{}

// Merge 接收多个structpb.Struct对象，合并它们并返回一个新的structpb.Struct对象。
func (m merger) Merge(values ...*structpb.Struct) *structpb.Struct {
	target := errorx.Ignore(structpb.NewStruct(map[string]any{}))
	for _, value := range values {
		m.mergeStruct(target, value)
	}
	return target
}

// mergeStruct 递归地合并两个structpb.Struct对象的字段。
func (m merger) mergeStruct(target *structpb.Struct, source *structpb.Struct) {
	for key, field := range source.GetFields() {
		target.Fields[key] = m.copyValue(field)
	}
}

// 递归地合并两个structpb.ListValue对象的值。
func (m merger) mergeList(target *structpb.ListValue, source *structpb.ListValue) {
	for _, item := range source.GetValues() {
		target.Values = append(target.Values, m.copyValue(item))
	}
}

// copyValue 根据structpb.Value的类型，创建并返回一个新的structpb.Value对象。
func (m merger) copyValue(value *structpb.Value) *structpb.Value {
	if value == nil {
		return structpb.NewNullValue()
	}
	switch v := value.GetKind().(type) {
	case *structpb.Value_NumberValue:
		if v == nil {
			return structpb.NewNullValue()
		}
		return structpb.NewNumberValue(v.NumberValue)
	case *structpb.Value_StringValue:
		if v == nil {
			return structpb.NewNullValue()
		}
		return structpb.NewStringValue(v.StringValue)
	case *structpb.Value_BoolValue:
		if v == nil {
			return structpb.NewNullValue()
		}
		return structpb.NewBoolValue(v.BoolValue)
	case *structpb.Value_StructValue:
		if v == nil {
			return structpb.NewNullValue()
		}
		subValue := errorx.Ignore(structpb.NewStruct(map[string]any{}))
		m.mergeStruct(subValue, v.StructValue)
		return structpb.NewStructValue(subValue)
	case *structpb.Value_ListValue:
		if v == nil {
			return structpb.NewNullValue()
		}
		subList := errorx.Ignore(structpb.NewList([]any{}))
		m.mergeList(subList, v.ListValue)
		return structpb.NewListValue(subList)
	case *structpb.Value_NullValue:
		return structpb.NewNullValue()
	default:
		return structpb.NewNullValue()
	}
}
