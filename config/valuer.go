package config

import (
	"errors"

	"github.com/go-leo/gox/stringx"
)

var ErrValueNotFound = errors.New("value not found")

var ErrValueIsNil = errors.New("value is nil")

const allConfigKey = ""

// valuer is Gets the value of a key
type valuer struct {
	data *Data
}

func (valuer *valuer) Value(key string) *Value {
	if stringx.IsBlank(key) {
		return &Value{val: valuer.data.AsMap()}
	}
	node, ok := valuer.data.AsTree().Find(key)
	if !ok {
		return &Value{err: ErrValueNotFound}
	}
	meta := node.Meta()
	if meta == nil {
		return &Value{err: ErrValueIsNil}
	}
	return &Value{val: meta}
}
