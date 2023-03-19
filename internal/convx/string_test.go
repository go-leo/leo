package convx_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-leo/gox/convx"
)

func TestString(t *testing.T) {
	tests := []struct {
		Arg any
	}{
		{
			Arg: true,
		},
		{
			Arg: int(10),
		},
		{
			Arg: int8(10),
		},
		{
			Arg: int16(10),
		},
		{
			Arg: int32(10),
		},
		{
			Arg: int64(10),
		},
		{
			Arg: uint(10),
		},
		{
			Arg: uint8(10),
		},
		{
			Arg: uint16(10),
		},
		{
			Arg: uint32(10),
		},
		{
			Arg: uint64(10),
		},
		{
			Arg: uintptr(10),
		},
		{
			Arg: float32(10),
		},
		{
			Arg: float64(10),
		},
		{
			Arg: "10",
		},
		{
			Arg: []byte("10"),
		},
		{
			Arg: SS{Num: 101},
		},
		{
			Arg: &SS{Num: 102},
		},
		// {
		// 	Arg: PS{Num: 103},
		// },
		{
			Arg: &PS{Num: 104},
		},
		{
			Arg: json.Number("json.Number"),
		},
	}

	for _, test := range tests {
		i, err := convx.String(test.Arg)
		assert.NoError(t, err)
		t.Log(i)
	}

}

type SS struct {
	Num int
}

func (s SS) String() string {
	return strconv.Itoa(s.Num)
}

type PS struct {
	Num int
}

func (s *PS) String() string {
	return strconv.Itoa(s.Num)
}
