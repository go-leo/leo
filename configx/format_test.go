package configx

import (
	"fmt"
	"github.com/go-leo/leo/v3/configx/test"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestENV_Support(t *testing.T) {
	p := &Env{} // replace with the actual package name where Env is defined

	// Test case for supported format
	format := &MockFormatter{format: "env"}
	assert.True(t, p.Support(format), "Expected format 'env' to be supported")

	// Test case for unsupported format
	format = &MockFormatter{format: "json"}
	assert.False(t, p.Support(format), "Expected format 'json' to be unsupported")
}

func TestENV_Parse(t *testing.T) {
	p := &Env{} // replace with the actual package name where Env is defined
	data := []byte("KEY1=value1\nKEY2=value2")

	// Successful parsing test case
	expected := &structpb.Struct{Fields: map[string]*structpb.Value{
		"KEY1": {Kind: &structpb.Value_StringValue{StringValue: "value1"}},
		"KEY2": {Kind: &structpb.Value_StringValue{StringValue: "value2"}},
	}}
	result, err := p.Parse(data)
	assert.NoError(t, err)
	assert.EqualValuesf(t, expected.AsMap(), result.AsMap(), "Parsed struct should match the expected value")
}

func TestJSON_Support(t *testing.T) {
	json := &Json{}
	format := &MockFormatter{format: "json"} // You might need to create a mock or stub for this
	assert.True(t, json.Support(format), "Expected Json to support 'json' format")
}

func TestJSON_Parse(t *testing.T) {
	json := &Json{}
	data := []byte(`{"field1": "value1"}`)
	result, err := json.Parse(data)
	assert.NoError(t, err, "Parsing should not produce an error")
	assert.IsType(t, &structpb.Struct{}, result, "Expected result to be of type *structpb.Struct")
	assert.EqualValues(t, "value1", result.Fields["field1"].GetStringValue(), "Expected field value to be 'value1'")
}

func TestTOML_Support(t *testing.T) {
	p := &Toml{}
	// 测试支持的格式
	assert.True(t, p.Support(&MockFormatter{"toml"}))
}

func TestTOML_Parse(t *testing.T) {
	p := &Toml{}
	data := []byte(`
title = "Toml Example"
owner.name = "Tom Preston-Werner"
owner.dob = 1979-05-27T07:32:00Z
[database]
server = "192.168.1.1"
ports = [ 8000, 8001, 8002 ]
connection_max = 5000
enabled = true
`)

	// 正确的源应当返回一个空的Structpb.Struct和nil错误
	result, err := p.Parse(data)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 测试结果是否符合预期，这里只是基本的非空和类型检查
	// 详细的字段检查应该在解析后的struct上进行
	assert.IsType(t, &structpb.Struct{}, result)
	slice := result.Fields["database"].GetStructValue().GetFields()["ports"].GetListValue().AsSlice()
	assert.EqualValues(t, []any{float64(8000), float64(8001), float64(8002)}, slice, "Expected field value to be 'value1'")

}

func TestYAML_Support(t *testing.T) {
	yamlParser := &Yaml{}

	// Test for supported formats
	supportedFormats := []Formatter{
		&MockFormatter{format: "yaml"},
		&MockFormatter{format: "yml"},
	}
	for _, format := range supportedFormats {
		assert.True(t, yamlParser.Support(format), "Expected format to be supported")
	}

	// Test for unsupported formats
	unsupportedFormats := []Formatter{
		&MockFormatter{format: "json"},
		&MockFormatter{format: "xml"},
		&MockFormatter{format: "txt"},
	}
	for _, format := range unsupportedFormats {
		assert.False(t, yamlParser.Support(format), "Expected format to be unsupported")
	}
}

func TestYAML_Parse(t *testing.T) {
	yamlParser := &Yaml{}
	validYAML := []byte("key: value")

	result, err := yamlParser.Parse(validYAML)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 测试结果是否符合预期，这里只是基本的非空和类型检查
	// 详细的字段检查应该在解析后的struct上进行
	assert.IsType(t, &structpb.Struct{}, result)
	assert.EqualValues(t, "value", result.GetFields()["key"].GetStringValue(), "Expected field value to be 'value1'")

}

func TestProto_Support(t *testing.T) {
	parser := &Proto{}

	// Test for supported formats
	supportedFormats := []Formatter{
		&MockFormatter{format: "proto"},
		&MockFormatter{format: "pb"},
	}
	for _, format := range supportedFormats {
		assert.True(t, parser.Support(format), "Expected format to be supported")
	}

	// Test for unsupported formats
	unsupportedFormats := []Formatter{
		&MockFormatter{format: "protobuf"},
		&MockFormatter{format: "protobuffer"},
	}
	for _, format := range unsupportedFormats {
		assert.False(t, parser.Support(format), "Expected format to be unsupported")
	}
}

func TestProto_Parse(t *testing.T) {
	application := test.Application{
		Addr: "localhost",
		Port: 8080,
	}
	validProto, _ := proto.Marshal(&application)

	protoParser := &Proto{Message: &test.Application{}}
	result, err := protoParser.Parse(validProto)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// 测试结果是否符合预期，这里只是基本的非空和类型检查
	// 详细的字段检查应该在解析后的struct上进行
	assert.IsType(t, &structpb.Struct{}, result)
	assert.EqualValues(t, "localhost", result.GetFields()["addr"].GetStringValue(), "Expected field value to be 'localhost'")
}

// MockFormatter is a mock implementation of the Formatter interface for testing purposes
type MockFormatter struct {
	format string
}

func (m *MockFormatter) Format() string {
	return m.format
}

func TestWrongYaml(t *testing.T) {
	c := `redis:
    addr: localhost:6379
    db: 0
    network: tcp
    password: test`
	var conf map[string]any
	err := yaml.Unmarshal([]byte(c), &conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	data, err := yaml.Marshal(map[string]any{
		"redis": map[string]any{
			"network":  "tcp",
			"addr":     "localhost:6379",
			"password": "test",
			"db":       0,
		},
	})
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(string(data))

}
