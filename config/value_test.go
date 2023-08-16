package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"html/template"
	"reflect"
	"testing"
	"time"
)

func TestTimeInUTC(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect time.Time
		iserr  bool
	}{
		{"2009-11-10 23:00:00 +0000 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // Time.String()
		{"Tue Nov 10 23:00:00 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},        // ANSIC
		{"Tue Nov 10 23:00:00 UTC 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},    // UnixDate
		{"Tue Nov 10 23:00:00 +0000 2009", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},  // RubyDate
		{"10 Nov 09 23:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},             // RFC822
		{"10 Nov 09 23:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},           // RFC822Z
		{"Tuesday, 10-Nov-09 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC850
		{"Tue, 10 Nov 2009 23:00:00 UTC", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},   // RFC1123
		{"Tue, 10 Nov 2009 23:00:00 +0000", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false}, // RFC1123Z
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339
		{"2018-10-21T23:21:29+0200", time.Date(2018, 10, 21, 21, 21, 29, 0, time.UTC), false},      // RFC3339 without timezone hh:mm colon
		{"2009-11-10T23:00:00Z", time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC), false},            // RFC3339Nano
		{"11:00PM", time.Date(0, 1, 1, 23, 0, 0, 0, time.UTC), false},                              // Kitchen
		{"Nov 10 23:00:00", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                    // Stamp
		{"Nov 10 23:00:00.000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},                // StampMilli
		{"Nov 10 23:00:00.000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},             // StampMicro
		{"Nov 10 23:00:00.000000000", time.Date(0, 11, 10, 23, 0, 0, 0, time.UTC), false},          // StampNano
		{"2016-03-06 15:28:01-00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},        // RFC3339 without T
		{"2016-03-06 15:28:01-0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},         // RFC3339 without T or timezone hh:mm colon
		{"2016-03-06 15:28:01", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -0000", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 -00:00", time.Date(2016, 3, 6, 15, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 +0900", time.Date(2016, 3, 6, 6, 28, 1, 0, time.UTC), false},
		{"2016-03-06 15:28:01 +09:00", time.Date(2016, 3, 6, 6, 28, 1, 0, time.UTC), false},
		{"2006-01-02", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{"02 Jan 2006", time.Date(2006, 1, 2, 0, 0, 0, 0, time.UTC), false},
		{1472574600, time.Date(2016, 8, 30, 16, 30, 0, 0, time.UTC), false},
		{int(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{int64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{int32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint(1482597504), time.Date(2016, 12, 24, 16, 38, 24, 0, time.UTC), false},
		{uint64(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{uint32(1234567890), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		{time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), time.Date(2009, 2, 13, 23, 31, 30, 0, time.UTC), false},
		// errors
		{"2006", time.Time{}, true},
		{testing.T{}, time.Time{}, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.TimeInUTC()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v.UTC(), test.expect)
	}
}

func TestDuration(t *testing.T) {

	var td time.Duration = 5
	tests := []struct {
		input  interface{}
		expect time.Duration
		iserr  bool
	}{
		{time.Duration(5), td, false},
		{int(5), td, false},
		{int64(5), td, false},
		{int32(5), td, false},
		{int16(5), td, false},
		{int8(5), td, false},
		{uint(5), td, false},
		{uint64(5), td, false},
		{uint32(5), td, false},
		{uint16(5), td, false},
		{uint8(5), td, false},
		{float64(5), td, false},
		{float32(5), td, false},
		{string("5"), td, false},
		{string("5ns"), td, false},
		{string("5us"), time.Microsecond * td, false},
		{string("5µs"), time.Microsecond * td, false},
		{string("5ms"), time.Millisecond * td, false},
		{string("5s"), time.Second * td, false},
		{string("5m"), time.Minute * td, false},
		{string("5h"), time.Hour * td, false},
		// errors
		{"test", 0, true},
		{testing.T{}, 0, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Duration()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect bool
		iserr  bool
	}{
		{0, false, false},
		{int64(0), false, false},
		{int32(0), false, false},
		{int16(0), false, false},
		{int8(0), false, false},
		{uint(0), false, false},
		{uint64(0), false, false},
		{uint32(0), false, false},
		{uint16(0), false, false},
		{uint8(0), false, false},
		{float64(0), false, false},
		{float32(0), false, false},
		{time.Duration(0), false, false},
		{nil, false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"False", false, false},
		{"f", false, false},
		{"F", false, false},
		{false, false, false},

		{"true", true, false},
		{"TRUE", true, false},
		{"True", true, false},
		{"t", true, false},
		{"T", true, false},
		{1, true, false},
		{int64(1), true, false},
		{int32(1), true, false},
		{int16(1), true, false},
		{int8(1), true, false},
		{uint(1), true, false},
		{uint64(1), true, false},
		{uint32(1), true, false},
		{uint16(1), true, false},
		{uint8(1), true, false},
		{float64(1), true, false},
		{float32(1), true, false},
		{time.Duration(1), true, false},
		{true, true, false},
		{-1, true, false},
		{int64(-1), true, false},
		{int32(-1), true, false},
		{int16(-1), true, false},
		{int8(-1), true, false},

		// errors
		{"test", false, true},
		{testing.T{}, false, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Bool()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

type testStep struct {
	input  interface{}
	expect interface{}
	iserr  bool
}

func createNumberTestSteps(zero, one, eight, eightnegative, eightpoint31, eightpoint31negative interface{}) []testStep {
	var jeight, jminuseight, jfloateight json.Number
	_ = json.Unmarshal([]byte("8"), &jeight)
	_ = json.Unmarshal([]byte("-8"), &jminuseight)
	_ = json.Unmarshal([]byte("8.0"), &jfloateight)

	kind := reflect.TypeOf(zero).Kind()
	isUint := kind == reflect.Uint || kind == reflect.Uint8 || kind == reflect.Uint16 || kind == reflect.Uint32 || kind == reflect.Uint64

	// Some precision is lost when converting from float64 to float32.
	eightpoint31_32 := eightpoint31
	eightpoint31negative_32 := eightpoint31negative
	if kind == reflect.Float64 {
		eightpoint31_32 = float64(float32(eightpoint31.(float64)))
		eightpoint31negative_32 = float64(float32(eightpoint31negative.(float64)))
	}

	return []testStep{
		{int(8), eight, false},
		{int8(8), eight, false},
		{int16(8), eight, false},
		{int32(8), eight, false},
		{int64(8), eight, false},
		{time.Weekday(8), eight, false},
		{time.Month(8), eight, false},
		{uint(8), eight, false},
		{uint8(8), eight, false},
		{uint16(8), eight, false},
		{uint32(8), eight, false},
		{uint64(8), eight, false},
		{float32(8.31), eightpoint31_32, false},
		{float64(8.31), eightpoint31, false},
		{true, one, false},
		{false, zero, false},
		{"8", eight, false},
		{nil, zero, false},
		{int(-8), eightnegative, isUint},
		{int8(-8), eightnegative, isUint},
		{int16(-8), eightnegative, isUint},
		{int32(-8), eightnegative, isUint},
		{int64(-8), eightnegative, isUint},
		{float32(-8.31), eightpoint31negative_32, isUint},
		{float64(-8.31), eightpoint31negative, isUint},
		{"-8", eightnegative, isUint},
		{jeight, eight, false},
		{jminuseight, eightnegative, isUint},
		{jfloateight, eight, false},
		{"test", zero, true},
		{testing.T{}, zero, true},
	}
}

func TestFloat64(t *testing.T) {
	tests := createNumberTestSteps(float64(0), float64(1), float64(8), float64(-8), float64(8.31), float64(-8.31))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Float64()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestFloat32(t *testing.T) {
	tests := createNumberTestSteps(float32(0), float32(1), float32(8), float32(-8), float32(8.31), float32(-8.31))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Float32()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestInt64(t *testing.T) {
	tests := createNumberTestSteps(int64(0), int64(1), int64(8), int64(-8), int64(8), int64(-8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int64()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestInt32(t *testing.T) {
	tests := createNumberTestSteps(int32(0), int32(1), int32(8), int32(-8), int32(8), int32(-8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int32()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestInt16(t *testing.T) {
	tests := createNumberTestSteps(int16(0), int16(1), int16(8), int16(-8), int16(8), int16(-8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int16()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestInt8(t *testing.T) {
	tests := createNumberTestSteps(int8(0), int8(1), int8(8), int8(-8), int8(8), int8(-8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int8()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestInt(t *testing.T) {
	tests := createNumberTestSteps(int(0), int(1), int(8), int(-8), int(8), int(-8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int()
		assert.Equal(t, err != nil, test.iserr)
		assert.Equal(t, v, test.expect)
	}
}

func TestUInt64(t *testing.T) {
	tests := createNumberTestSteps(uint64(0), uint64(1), uint64(8), uint64(0), uint64(8), uint64(8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.UInt64()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestUInt32(t *testing.T) {
	tests := createNumberTestSteps(uint32(0), uint32(1), uint32(8), uint32(0), uint32(8), uint32(8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.UInt32()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestUInt16(t *testing.T) {
	tests := createNumberTestSteps(uint16(0), uint16(1), uint16(8), uint16(0), uint16(8), uint16(8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.UInt16()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestUInt8(t *testing.T) {
	tests := createNumberTestSteps(uint8(0), uint8(1), uint8(8), uint8(0), uint8(8), uint8(8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.UInt8()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestUInt(t *testing.T) {
	tests := createNumberTestSteps(uint(0), uint(1), uint(8), uint(0), uint(8), uint(8))
	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.UInt()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestString(t *testing.T) {
	var jn json.Number
	_ = json.Unmarshal([]byte("8"), &jn)
	type Key struct {
		k string
	}
	key := &Key{"foo"}

	tests := []struct {
		input  interface{}
		expect string
		iserr  bool
	}{
		{int(8), "8", false},
		{int8(8), "8", false},
		{int16(8), "8", false},
		{int32(8), "8", false},
		{int64(8), "8", false},
		{uint(8), "8", false},
		{uint8(8), "8", false},
		{uint16(8), "8", false},
		{uint32(8), "8", false},
		{uint64(8), "8", false},
		{float32(8.31), "8.31", false},
		{float64(8.31), "8.31", false},
		{jn, "8", false},
		{true, "true", false},
		{false, "false", false},
		{nil, "", false},
		{[]byte("one time"), "one time", false},
		{"one more time", "one more time", false},
		{template.HTML("one time"), "one time", false},
		{template.URL("http://somehost.foo"), "http://somehost.foo", false},
		{template.JS("(1+2)"), "(1+2)", false},
		{template.CSS("a"), "a", false},
		{template.HTMLAttr("a"), "a", false},
		// errors
		{testing.T{}, "", true},
		{key, "", true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.String()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestSlice(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []interface{}
		iserr  bool
	}{
		{[]interface{}{1, 3}, []interface{}{1, 3}, false},
		{[]map[string]interface{}{{"k1": 1}, {"k2": 2}}, []interface{}{map[string]interface{}{"k1": 1}, map[string]interface{}{"k2": 2}}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Slice()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestToBoolSliceE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []bool
		iserr  bool
	}{
		{[]bool{true, false, true}, []bool{true, false, true}, false},
		{[]interface{}{true, false, true}, []bool{true, false, true}, false},
		{[]int{1, 0, 1}, []bool{true, false, true}, false},
		{[]string{"true", "false", "true"}, []bool{true, false, true}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.BoolSlice()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestStringSlice(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []string
		iserr  bool
	}{
		{[]int{1, 2}, []string{"1", "2"}, false},
		{[]int8{int8(1), int8(2)}, []string{"1", "2"}, false},
		{[]int32{int32(1), int32(2)}, []string{"1", "2"}, false},
		{[]int64{int64(1), int64(2)}, []string{"1", "2"}, false},
		{[]float32{float32(1.01), float32(2.01)}, []string{"1.01", "2.01"}, false},
		{[]float64{float64(1.01), float64(2.01)}, []string{"1.01", "2.01"}, false},
		{[]string{"a", "b"}, []string{"a", "b"}, false},
		{[]interface{}{1, 3}, []string{"1", "3"}, false},
		{interface{}(1), []string{"1"}, false},
		{[]error{errors.New("a"), errors.New("b")}, []string{"a", "b"}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.StringSlice()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestIntSlice(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []int
		iserr  bool
	}{
		{[]int{1, 3}, []int{1, 3}, false},
		{[]interface{}{1.2, 3.2}, []int{1, 3}, false},
		{[]string{"2", "3"}, []int{2, 3}, false},
		{[2]string{"2", "3"}, []int{2, 3}, false},
		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"foo", "bar"}, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.IntSlice()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestDurationSlice(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect []time.Duration
		iserr  bool
	}{
		{[]string{"1s", "1m"}, []time.Duration{time.Second, time.Minute}, false},
		{[]int{1, 2}, []time.Duration{1, 2}, false},
		{[]interface{}{1, 3}, []time.Duration{1, 3}, false},
		{[]time.Duration{1, 3}, []time.Duration{1, 3}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{[]string{"invalid"}, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.DurationSlice()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]interface{}
		iserr  bool
	}{
		{map[interface{}]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{map[string]interface{}{"tag": "tags", "group": "groups"}, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": "groups"}`, map[string]interface{}{"tag": "tags", "group": "groups"}, false},
		{`{"tag": "tags", "group": true}`, map[string]interface{}{"tag": "tags", "group": true}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Map()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestStringMap(t *testing.T) {
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var jsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}`
	var invalidJsonString = `{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"`
	var emptyString = ""

	tests := []struct {
		input  interface{}
		expect map[string]string
		iserr  bool
	}{
		{stringMapString, stringMapString, false},
		{stringMapInterface, stringMapString, false},
		{interfaceMapString, stringMapString, false},
		{interfaceMapInterface, stringMapString, false},
		{jsonString, stringMapString, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{invalidJsonString, nil, true},
		{emptyString, nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.StringMap()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestStringMapStringSliceE(t *testing.T) {
	// ToStringMapString inputs/outputs
	var stringMapString = map[string]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var stringMapInterface = map[string]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapString = map[interface{}]string{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}
	var interfaceMapInterface = map[interface{}]interface{}{"key 1": "value 1", "key 2": "value 2", "key 3": "value 3"}

	// ToStringMapStringSlice inputs/outputs
	var stringMapStringSlice = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceSlice = map[string][]interface{}{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapInterfaceInterfaceSlice = map[string]interface{}{"key 1": []interface{}{"value 1", "value 2", "value 3"}, "key 2": []interface{}{"value 1", "value 2", "value 3"}, "key 3": []interface{}{"value 1", "value 2", "value 3"}}
	var stringMapStringSingleSliceFieldsResult = map[string][]string{"key 1": {"value", "1"}, "key 2": {"value", "2"}, "key 3": {"value", "3"}}
	var interfaceMapStringSlice = map[interface{}][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var interfaceMapInterfaceSlice = map[interface{}][]interface{}{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}

	var stringMapStringSliceMultiple = map[string][]string{"key 1": {"value 1", "value 2", "value 3"}, "key 2": {"value 1", "value 2", "value 3"}, "key 3": {"value 1", "value 2", "value 3"}}
	var stringMapStringSliceSingle = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2"}, "key 3": {"value 3"}}

	var stringMapInterface1 = map[string]interface{}{"key 1": []string{"value 1"}, "key 2": []string{"value 2"}}
	var stringMapInterfaceResult1 = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2"}}

	var jsonStringMapString = `{"key 1": "value 1", "key 2": "value 2"}`
	var jsonStringMapStringArray = `{"key 1": ["value 1"], "key 2": ["value 2", "value 3"]}`
	var jsonStringMapStringArrayResult = map[string][]string{"key 1": {"value 1"}, "key 2": {"value 2", "value 3"}}

	type Key struct {
		k string
	}

	tests := []struct {
		input  interface{}
		expect map[string][]string
		iserr  bool
	}{
		{stringMapStringSlice, stringMapStringSlice, false},
		{stringMapInterfaceSlice, stringMapStringSlice, false},
		{stringMapInterfaceInterfaceSlice, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapStringSliceMultiple, stringMapStringSlice, false},
		{stringMapString, stringMapStringSliceSingle, false},
		{stringMapInterface, stringMapStringSliceSingle, false},
		{stringMapInterface1, stringMapInterfaceResult1, false},
		{interfaceMapStringSlice, stringMapStringSlice, false},
		{interfaceMapInterfaceSlice, stringMapStringSlice, false},
		{interfaceMapString, stringMapStringSingleSliceFieldsResult, false},
		{interfaceMapInterface, stringMapStringSingleSliceFieldsResult, false},
		{jsonStringMapStringArray, jsonStringMapStringArrayResult, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{map[interface{}]interface{}{"foo": testing.T{}}, nil, true},
		{map[interface{}]interface{}{Key{"foo"}: "bar"}, nil, true}, // ToStringE(Key{"foo"}) should fail
		{jsonStringMapString, nil, true},
		{"", nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.StringSliceMap()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestToStringMapIntE(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]int
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": 1, "v2": 222}, map[string]int{"v1": 1, "v2": 222}, false},
		{map[string]interface{}{"v1": 342, "v2": 5141}, map[string]int{"v1": 342, "v2": 5141}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.IntMap()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestInt64Map(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]int64
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": int32(8), "v2": int32(888)}, map[string]int64{"v1": int64(8), "v2": int64(888)}, false},
		{map[string]interface{}{"v1": int64(45), "v2": int64(67)}, map[string]int64{"v1": 45, "v2": 67}, false},
		{map[string]int64{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int{"v1": 33, "v2": 88}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]int32{"v1": int32(33), "v2": int32(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]uint16{"v1": uint16(33), "v2": uint16(88)}, map[string]int64{"v1": 33, "v2": 88}, false},
		{map[string]float64{"v1": float64(8.22), "v2": float64(43.32)}, map[string]int64{"v1": 8, "v2": 43}, false},
		{`{"v1": 67, "v2": 56}`, map[string]int64{"v1": 67, "v2": 56}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.Int64Map()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestBoolMap(t *testing.T) {
	tests := []struct {
		input  interface{}
		expect map[string]bool
		iserr  bool
	}{
		{map[interface{}]interface{}{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]interface{}{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{map[string]bool{"v1": true, "v2": false}, map[string]bool{"v1": true, "v2": false}, false},
		{`{"v1": true, "v2": false}`, map[string]bool{"v1": true, "v2": false}, false},

		// errors
		{nil, nil, true},
		{testing.T{}, nil, true},
		{"", nil, true},
	}

	for _, test := range tests {
		val := Value{val: test.input}
		v, err := val.BoolMap()
		assert.Equal(t, test.iserr, err != nil)
		if !test.iserr {
			assert.Equal(t, test.expect, v)
		}
	}
}

func TestScan(t *testing.T) {
	type Family struct {
		LastName string `mapstructure:"LastName" json:"last_name" yaml:"last_name"`
	}
	type Location struct {
		City string `mapstructure:"City" json:"city" yaml:"city"`
	}
	type Person struct {
		Family    `mapstructure:",squash" json:"family" yaml:"family"`
		Location  `mapstructure:",squash" json:"location" yaml:"location"`
		FirstName string            `mapstructure:"FirstName" json:"first_name" yaml:"first_name"`
		Name      string            `mapstructure:"Name" json:"name" yaml:"name"`
		Age       int               `mapstructure:"Age" json:"age" yaml:"age"`
		Emails    []string          `mapstructure:"Emails" json:"emails" yaml:"emails"`
		Extra     map[string]string `mapstructure:"Extra" json:"extra" yaml:"extra"`
	}

	input := map[string]interface{}{
		"FirstName": "Mitchell",
		"LastName":  "Hashimoto",
		"City":      "San Francisco",
		"name":      "Mitchell",
		"age":       91,
		"emails":    []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	val := Value{val: input}
	err := val.Scan(&result)
	assert.NoError(t, err)
	assert.Equal(t, Person{
		Family: Family{
			LastName: "Hashimoto",
		},
		Location: Location{
			City: "San Francisco",
		},
		FirstName: "Mitchell",
		Name:      "Mitchell",
		Age:       91,
		Emails:    []string{"one", "two", "three"},
		Extra: map[string]string{
			"twitter": "mitchellh",
		},
	}, result)

	data, _ := json.Marshal(&result)
	fmt.Println(string(data))
}
