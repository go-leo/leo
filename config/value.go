package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cast"
	"time"
)

type Value struct {
	val any
	err error
}

func (v Value) TimeInLocal() (time.Time, error) {
	if v.err != nil {
		return time.Time{}, v.err
	}
	return cast.ToTimeInDefaultLocationE(v.val, time.Local)
}

func (v Value) TimeInUTC() (time.Time, error) {
	if v.err != nil {
		return time.Time{}, v.err
	}
	return cast.ToTimeInDefaultLocationE(v.val, time.UTC)
}

func (v Value) Duration() (time.Duration, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToDurationE(v.val)
}

func (v Value) Bool() (bool, error) {
	if v.err != nil {
		return false, v.err
	}
	return cast.ToBoolE(v.val)
}

func (v Value) Float64() (float64, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToFloat64E(v.val)
}

func (v Value) Float32() (float32, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToFloat32E(v.val)
}

func (v Value) Int64() (int64, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToInt64E(v.val)
}

func (v Value) Int32() (int32, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToInt32E(v.val)
}

func (v Value) Int16() (int16, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToInt16E(v.val)
}

func (v Value) Int8() (int8, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToInt8E(v.val)
}

func (v Value) Int() (int, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToIntE(v.val)
}

func (v Value) UInt64() (uint64, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToUint64E(v.val)
}

func (v Value) UInt32() (uint32, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToUint32E(v.val)
}

func (v Value) UInt16() (uint16, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToUint16E(v.val)
}

func (v Value) UInt8() (uint8, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToUint8E(v.val)
}

func (v Value) UInt() (uint, error) {
	if v.err != nil {
		return 0, v.err
	}
	return cast.ToUintE(v.val)
}

func (v Value) String() (string, error) {
	if v.err != nil {
		return "", v.err
	}
	return cast.ToStringE(v.val)
}

func (v Value) Slice() ([]any, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToSliceE(v.val)
}

func (v Value) BoolSlice() ([]bool, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToBoolSliceE(v.val)
}

func (v Value) StringSlice() ([]string, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringSliceE(v.val)
}

func (v Value) IntSlice() ([]int, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToIntSliceE(v.val)
}

func (v Value) DurationSlice() ([]time.Duration, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToDurationSliceE(v.val)
}

func (v Value) Map() (map[string]any, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapE(v.val)
}

func (v Value) StringMap() (map[string]string, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapStringE(v.val)
}

func (v Value) StringSliceMap() (map[string][]string, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapStringSliceE(v.val)
}

func (v Value) IntMap() (map[string]int, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapIntE(v.val)
}

func (v Value) Int64Map() (map[string]int64, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapInt64E(v.val)
}

func (v Value) BoolMap() (map[string]bool, error) {
	if v.err != nil {
		return nil, v.err
	}
	return cast.ToStringMapBoolE(v.val)
}

func (v Value) Scan(output any) error {
	if v.err != nil {
		return v.err
	}
	config := DefaultMapStructureConfig
	config.Result = output
	decoder, err := mapstructure.NewDecoder(&config)
	if err != nil {
		return err
	}
	return decoder.Decode(v.val)
}

var DefaultMapStructureConfig = mapstructure.DecoderConfig{
	DecodeHook: mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToIPHookFunc(),
		mapstructure.StringToIPNetHookFunc(),
		mapstructure.WeaklyTypedHook,
		mapstructure.RecursiveStructToMapHookFunc(),
		mapstructure.TextUnmarshallerHookFunc()),
	ErrorUnused:          false,
	ErrorUnset:           false,
	ZeroFields:           false,
	WeaklyTypedInput:     false,
	Squash:               false,
	Metadata:             nil,
	Result:               nil,
	TagName:              "",
	IgnoreUntaggedFields: false,
	MatchName:            nil,
}
