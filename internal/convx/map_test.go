package convx_test

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/gox/convx"
)

func TestMap(t *testing.T) {
	tests := []struct {
		Arg any
	}{
		{
			Arg: map[string]int{"num": 10},
		},
		{
			Arg: map[int]int8{1: 10},
		},
		{
			Arg: map[bool]int16{true: 10},
		},
		{
			Arg: map[float64]int32{1.1: 10},
		},
		{
			Arg: map[uint32]int64{0: 10},
		},
		{
			Arg: map[uint64]uint{10: 10},
		},
		{
			Arg: map[io.Reader]uint8{io.NopCloser(nil): 10},
		},
		{
			Arg: map[complex64]uint16{complex64(10i): 10},
		},
		{
			Arg: map[[10]int]uint32{[10]int{1}: 10},
		},
		{
			Arg: map[float32]uint64{2.3: 10},
		},
		{
			Arg: map[string]uintptr{"ff": 10},
		},
		{
			Arg: map[int]float32{50: 10},
		},
		{
			Arg: map[float64]float64{10: 11.1},
		},
		{
			Arg: map[bool]string{false: "10"},
		},
		{
			Arg: map[struct{ Key int }]byte{struct{ Key int }{Key: 10}: 10},
		},
	}

	for _, test := range tests {
		i, err := convx.Map(test.Arg)
		assert.NoError(t, err)
		t.Log(i)
	}
}
