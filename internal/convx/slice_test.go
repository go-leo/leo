package convx_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/gox/convx"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		Arg any
	}{
		{
			Arg: []int{10},
		},
		{
			Arg: []int8{10},
		},
		{
			Arg: []int16{10},
		},
		{
			Arg: []int32{10},
		},
		{
			Arg: []int64{10},
		},
		{
			Arg: []uint{10},
		},
		{
			Arg: []uint8{10},
		},
		{
			Arg: []uint16{10},
		},
		{
			Arg: []uint32{10},
		},
		{
			Arg: []uint64{10},
		},
		{
			Arg: []uintptr{10},
		},
		{
			Arg: []float32{10},
		},
		{
			Arg: []float64{10},
		},
		{
			Arg: []string{"10"},
		},
		{
			Arg: []byte("10"),
		},
	}

	for _, test := range tests {
		i, err := convx.Slice(test.Arg)
		assert.NoError(t, err)
		t.Log(i)
	}
}
