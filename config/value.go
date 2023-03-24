package config

import (
	"time"

	"github.com/go-leo/gox/convx"
	"github.com/go-leo/gox/encodingx/mapstructure"
)

type Value struct {
	val any
}

func (v Value) TimeInLocal() (time.Time, error) {
	return convx.ToTimeInDefaultLocationE(v.val, time.Local)
}

func (v Value) TimeInUTC() (time.Time, error) {
	return convx.ToTimeInDefaultLocationE(v.val, time.UTC)
}

func (v Value) Duration() (time.Duration, error) {
	return convx.ToDurationE(v.val)
}

func (v Value) Bool() (bool, error) {
	return convx.ToBoolE(v.val)
}

func (v Value) Float64() (float64, error) {
	return convx.ToFloat64E(v.val)
}

func (v Value) Float32() (float32, error) {
	return convx.ToFloat32E(v.val)
}

func (v Value) Int64() (int64, error) {
	return convx.ToInt64E(v.val)
}

func (v Value) Int32() (int32, error) {
	return convx.ToInt32E(v.val)
}

func (v Value) Int16() (int16, error) {
	return convx.ToInt16E(v.val)
}

func (v Value) Int8() (int8, error) {
	return convx.ToInt8E(v.val)
}

func (v Value) Int() (int, error) {
	return convx.ToIntE(v.val)
}

func (v Value) UInt64() (uint64, error) {
	return convx.ToUint64E(v.val)
}

func (v Value) UInt32() (uint32, error) {
	return convx.ToUint32E(v.val)
}

func (v Value) UInt16() (uint16, error) {
	return convx.ToUint16E(v.val)
}

func (v Value) UInt8() (uint8, error) {
	return convx.ToUint8E(v.val)
}

func (v Value) UInt() (uint, error) {
	return convx.ToUintE(v.val)
}

func (v Value) String() (string, error) {
	return convx.ToStringE(v.val)
}

func (v Value) Slice() ([]any, error) {
	return convx.ToSliceE(v.val)
}

func (v Value) BoolSlice() ([]bool, error) {
	return convx.ToBoolSliceE(v.val)
}

func (v Value) StringSlice() ([]string, error) {
	return convx.ToStringSliceE(v.val)
}

func (v Value) IntSlice() ([]int, error) {
	return convx.ToIntSliceE(v.val)
}

func (v Value) DurationSlice() ([]time.Duration, error) {
	return convx.ToDurationSliceE(v.val)
}

func (v Value) Map() (map[string]any, error) {
	return convx.ToStringMapE(v.val)
}

func (v Value) StringMap() (map[string]string, error) {
	return convx.ToStringMapStringE(v.val)
}

func (v Value) StringSliceMap() (map[string][]string, error) {
	return convx.ToStringMapStringSliceE(v.val)
}

func (v Value) IntMap() (map[string]int, error) {
	return convx.ToStringMapIntE(v.val)
}

func (v Value) Int64Map() (map[string]int64, error) {
	return convx.ToStringMapInt64E(v.val)
}

func (v Value) BoolMap() (map[string]bool, error) {
	return convx.ToStringMapBoolE(v.val)
}

func (v Value) MapStructure(output any) error {
	c := &mapstructure.DecoderConfig{
		Result:           output,
		WeaklyTypedInput: true,
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			mapstructure.StringToSliceHookFunc(","),
			mapstructure.StringToTimeDurationHookFunc(),
			mapstructure.StringToIPHookFunc(),
			mapstructure.StringToIPNetHookFunc(),
			mapstructure.StringToTimeHookFunc(time.RFC3339),
			mapstructure.WeaklyTypedHook,
			mapstructure.TextUnmarshallerHookFunc(),
		),
	}
	decoder, err := mapstructure.NewDecoder(c)
	if err != nil {
		return err
	}
	return decoder.Decode(v.val)
}
