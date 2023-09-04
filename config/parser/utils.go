package parser

import (
	"reflect"

	"github.com/spf13/cast"

	"github.com/hmldd/leo/config"
)

func Parsers() []config.Parser {
	return []config.Parser{
		NewYamlParser(),
		NewJsonParser(),
		NewTomlParser(),
	}
}

func Parser(contentType string) config.Parser {
	switch contentType {
	case YAML, YML:
		return NewYamlParser()
	case JSON:
		return NewJsonParser()
	case TOML:
		return NewTomlParser()
	default:
		return &NoopParser{}
	}
}

func standardized(arg any) any {
	if arg == nil {
		return arg
	}
	argType := reflect.TypeOf(arg)
	argValue := reflect.ValueOf(arg)
	switch argType.Kind() {
	case reflect.Slice:
		length := argValue.Len()
		slice := make([]any, 0, length)
		for i := 0; i < length; i++ {
			v := argValue.Index(i).Interface()
			slice = append(slice, standardized(v))
		}
		return slice
	case reflect.Map:
		length := argValue.Len()
		m := make(map[string]any, length)
		keys := argValue.MapKeys()
		for _, key := range keys {
			v := argValue.MapIndex(key).Interface()
			k := cast.ToString(key.Interface())
			m[k] = standardized(v)
		}
		return m
	default:
		return arg
	}
}
