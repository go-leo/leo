package configx

import (
	"github.com/go-leo/gox/errorx"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
)

func TestMerger_Merge(t *testing.T) {
	m := merger{}

	// 测试空值
	emptyStruct := m.Merge()
	if emptyStruct == nil || len(emptyStruct.GetFields()) != 0 {
		t.Errorf("Expected empty struct, got %v", emptyStruct)
	}

	// 测试单个结构合并
	singleStruct := m.Merge(errorx.Ignore(structpb.NewStruct(map[string]any{"a": 1})))
	if singleStruct == nil || len(singleStruct.GetFields()) != 1 || singleStruct.GetFields()["a"].GetKind().(*structpb.Value_NumberValue).NumberValue != 1 {
		t.Errorf("Expected single field with value 1, got %v", singleStruct)
	}

	// 测试多个结构合并
	multipleStructs := m.Merge(
		errorx.Ignore(structpb.NewStruct(map[string]any{"a": 1, "b": "one"})),
		errorx.Ignore(structpb.NewStruct(map[string]any{"b": "two", "c": true})),
	)
	if multipleStructs == nil || len(multipleStructs.GetFields()) != 3 {
		t.Errorf("Expected three fields, got %v", multipleStructs)
	}
	if multipleStructs.GetFields()["a"].GetKind().(*structpb.Value_NumberValue).NumberValue != 1 {
		t.Errorf("Field 'a' expected value 1, got %v", multipleStructs.GetFields()["a"])
	}
	if multipleStructs.GetFields()["b"].GetKind().(*structpb.Value_StringValue).StringValue != "two" {
		t.Errorf("Field 'b' expected value 'two', got %v", multipleStructs.GetFields()["b"])
	}
	if multipleStructs.GetFields()["c"].GetKind().(*structpb.Value_BoolValue).BoolValue != true {
		t.Errorf("Field 'c' expected value true, got %v", multipleStructs.GetFields()["c"])
	}
}
