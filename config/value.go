package config

import (
	"time"

	"github.com/spf13/cast"
)

type Value struct {
	val any
}

func (v Value) TimeInLocal() (time.Time, error) {
	return cast.ToTimeInDefaultLocationE(v.val, time.Local)
}

func (v Value) TimeInUTC() (time.Time, error) {
	return cast.ToTimeInDefaultLocationE(v.val, time.UTC)
}

func (v Value) Duration() (time.Duration, error) {
	return cast.ToDurationE(v.val)
}

func (v Value) Int() (int, error) {
	return cast.ToIntE(v.val)
}

func (v Value) Int8() (int8, error) {
	return cast.ToInt8E(v.val)
}

func (v Value) Int16() (int16, error) {
	return cast.ToInt16E(v.val)
}

func (v Value) Int32() (int32, error) {
	return cast.ToInt32E(v.val)
}

func (v Value) Int64() (int64, error) {
	return cast.ToInt64E(v.val)
}

func (v Value) UInt() (uint, error) {
	return cast.ToUintE(v.val)
}

func (v Value) UInt8() (uint8, error) {
	return cast.ToUint8E(v.val)
}

func (v Value) UInt16() (uint16, error) {
	return cast.ToUint16E(v.val)
}

func (v Value) UInt32() (uint32, error) {
	return cast.ToUint32E(v.val)
}

func (v Value) UInt64() (uint64, error) {
	return cast.ToUint64E(v.val)
}

func (v Value) Float32() (float32, error) {
	return cast.ToFloat32E(v.val)
}

func (v Value) Float64() (float64, error) {
	return cast.ToFloat64E(v.val)
}

func (v Value) Bool() (bool, error) {
	return cast.ToBoolE(v.val)
}

func (v Value) String() (string, error) {
	return cast.ToStringE(v.val)
}

func (v Value) Slice() ([]any, error) {
	return cast.ToSliceE(v.val)
}

func (v Value) BoolSlice() ([]bool, error) {
	return cast.ToBoolSliceE(v.val)
}

func (v Value) StringSlice() ([]string, error) {
	return cast.ToStringSliceE(v.val)
}

func (v Value) IntSlice() ([]int, error) {
	return cast.ToIntSliceE(v.val)
}

func (v Value) DurationSlice() ([]time.Duration, error) {
	return cast.ToDurationSliceE(v.val)
}

func (v Value) Map() (map[string]any, error) {
	return cast.ToStringMapE(v.val)
}

func (v Value) StringMap() (map[string]string, error) {
	return cast.ToStringMapStringE(v.val)
}

func (v Value) StringSliceMap() (map[string][]string, error) {
	return cast.ToStringMapStringSliceE(v.val)
}

func (v Value) IntMap() (map[string]int, error) {
	return cast.ToStringMapIntE(v.val)
}

func (v Value) Int64Map() (map[string]int64, error) {
	return cast.ToStringMapInt64E(v.val)
}

func (v Value) BoolMap() (map[string]bool, error) {
	return cast.ToStringMapBoolE(v.val)
}

func (v Value) Scan(out any, decode func(in any, out any) error) error {
	return decode(v.val, out)
}
