package config

import (
	"errors"

	"github.com/go-leo/gox/stringx"
)

var ValueNotFound = errors.New("value not found")

var Nil = errors.New("value is nil")

const allConfigKey = ""

// valuer is Gets the value of a key
type valuer struct {
	data *Data
}

func (valuer *valuer) Value(key string) (*Value, error) {
	if stringx.IsBlank(key) {
		return &Value{val: valuer.data.AsMap()}, nil
	}
	node, ok := valuer.data.AsTree().Find(key)
	if !ok {
		return nil, ValueNotFound
	}
	meta := node.Meta()
	if meta == nil {
		return nil, Nil
	}
	return &Value{val: meta}, nil
}
