package configx

import (
	"github.com/go-leo/gox/encodingx/envx"
	"github.com/go-leo/gox/encodingx/jsonx"
	"github.com/go-leo/gox/encodingx/tomlx"
	"github.com/go-leo/gox/encodingx/yamlx"
	"google.golang.org/protobuf/types/known/structpb"
	"slices"
	"strings"
)

var _ Formatter = (*Env)(nil)
var _ Parser = (*Env)(nil)

// Env Formatter
type Env struct{}

func (Env) Format() string { return "env" }

func (Env) Support(format Formatter) bool { return strings.EqualFold(format.Format(), "env") }

func (Env) Parse(data []byte) (*structpb.Struct, error) {
	v := make(map[string]string)
	if err := envx.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	jsonData, err := jsonx.Marshal(v)
	if err != nil {
		return nil, err
	}
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(jsonData)
}

var _ Formatter = (*Json)(nil)
var _ Parser = (*Json)(nil)

// Json Formatter
type Json struct{}

func (Json) Format() string { return "json" }

func (Json) Support(format Formatter) bool { return strings.EqualFold(format.Format(), "json") }
func (Json) Parse(data []byte) (*structpb.Struct, error) {
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(data)
}

var _ Formatter = (*Toml)(nil)
var _ Parser = (*Toml)(nil)

// Toml Formatter
type Toml struct{}

func (Toml) Format() string { return "toml" }

func (Toml) Support(format Formatter) bool { return strings.EqualFold(format.Format(), "toml") }

func (Toml) Parse(data []byte) (*structpb.Struct, error) {
	v := make(map[string]any)
	if err := tomlx.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	jsonData, err := jsonx.Marshal(v)
	if err != nil {
		return nil, err
	}
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(jsonData)
}

var _ Formatter = (*Yaml)(nil)
var _ Parser = (*Yaml)(nil)

// Yaml Formatter
type Yaml struct{}

func (Yaml) Format() string { return "yaml" }

func (Yaml) Support(format Formatter) bool {
	return slices.ContainsFunc([]string{"yaml", "yml"}, func(s string) bool { return strings.EqualFold(format.Format(), s) })
}

func (Yaml) Parse(data []byte) (*structpb.Struct, error) {
	v := make(map[string]any)
	if err := yamlx.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	jsonData, err := jsonx.Marshal(v)
	if err != nil {
		return nil, err
	}
	value := &structpb.Struct{Fields: make(map[string]*structpb.Value)}
	return value, value.UnmarshalJSON(jsonData)
}
