package convx_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/reflectx"
)

func TestInt(t *testing.T) {
	tests := []struct {
		Arg  any
		Kind reflect.Kind
	}{
		{
			Arg:  int(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  int8(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  int16(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  int32(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  int64(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uint(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uint8(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uint16(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uint32(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uint64(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  uintptr(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  float32(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  float64(10),
			Kind: reflect.Int64,
		},
		{
			Arg:  "10",
			Kind: reflect.Int64,
		},
		{
			Arg:  []byte("10"),
			Kind: reflect.Int64,
		},
	}

	for _, test := range tests {
		i, err := convx.Int(test.Arg)
		assert.NoError(t, err)
		kind := reflectx.BaseKind(i)
		assert.Equal(t, test.Kind, kind)
	}

}
