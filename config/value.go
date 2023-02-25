package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

var Nil = errors.New("value is nil")

type Value struct {
	val any
}

func (v Value) Int() (int64, error) {
	return toInt64(v.val)
}

func (v Value) Float() (float64, error) {
	return toFloat64(v.val)
}

func (v Value) Bool() (bool, error) {
	return toBool(v.val)
}

func (v Value) String() (string, error) {
	return toString(v.val)
}

func (v Value) Slice() ([]any, error) {
	return toSlice(v.val)
}

func (v Value) Map() (map[string]any, error) {
	a := v.val
	return toMap(a)
}

func (v Value) Scan(out any, decode func(in any, out any) error) error {
	return decode(v.val, out)
}

func toInt64(val any) (int64, error) {
	val = reflectx.Indirect(val)
	switch i := val.(type) {
	case int:
		return int64(i), nil
	case int64:
		return i, nil
	case int32:
		return int64(i), nil
	case int16:
		return int64(i), nil
	case int8:
		return int64(i), nil
	case uint:
		return int64(i), nil
	case uint64:
		return int64(i), nil
	case uint32:
		return int64(i), nil
	case uint16:
		return int64(i), nil
	case uint8:
		return int64(i), nil
	case float64:
		return int64(i), nil
	case float32:
		return int64(i), nil
	case string:
		v, err := strconv.ParseInt(trimZeroDecimal(i), 0, 0)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", val, val)
	case json.Number:
		return Value{val: string(i)}.Int()
	case bool:
		if i {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, Nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", val, val)
	}
}

func toFloat64(val any) (float64, error) {
	val = reflectx.Indirect(val)
	switch f := val.(type) {
	case int:
		return float64(f), nil
	case int64:
		return float64(f), nil
	case int32:
		return float64(f), nil
	case int16:
		return float64(f), nil
	case int8:
		return float64(f), nil
	case uint:
		return float64(f), nil
	case uint64:
		return float64(f), nil
	case uint32:
		return float64(f), nil
	case uint16:
		return float64(f), nil
	case uint8:
		return float64(f), nil
	case float64:
		return f, nil
	case float32:
		return float64(f), nil
	case string:
		v, err := strconv.ParseFloat(f, 64)
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", val, val)
	case json.Number:
		v, err := Value{val: string(f)}.Float()
		if err == nil {
			return v, nil
		}
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", val, val)
	case bool:
		if f {
			return 1, nil
		}
		return 0, nil
	case nil:
		return 0, Nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to float64", val, val)
	}
}

func toBool(val any) (bool, error) {
	val = reflectx.Indirect(val)
	switch b := val.(type) {
	case bool:
		return b, nil
	case string:
		return strconv.ParseBool(val.(string))
	case nil:
		return false, Nil
	default:
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", val, val)
	}
}

func toString(val any) (string, error) {
	val = reflectx.IndirectToStringerOrError(val)
	switch s := val.(type) {
	case string:
		return s, nil
	case bool:
		return strconv.FormatBool(s), nil
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(s), nil
	case int64:
		return strconv.FormatInt(s, 10), nil
	case int32:
		return strconv.Itoa(int(s)), nil
	case int16:
		return strconv.FormatInt(int64(s), 10), nil
	case int8:
		return strconv.FormatInt(int64(s), 10), nil
	case uint:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(s), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(s), 10), nil
	case json.Number:
		return s.String(), nil
	case fmt.Stringer:
		return s.String(), nil
	case []byte:
		return string(s), nil
	case error:
		return s.Error(), nil
	case nil:
		return "", Nil
	default:
		return "", fmt.Errorf("unable to cast %#v of type %T to string", val, val)
	}
}

func toSlice(val any) ([]any, error) {
	var r []any
	switch s := val.(type) {
	case []any:
		return append(r, s...), nil
	case []map[string]any:
		for _, u := range s {
			r = append(r, u)
		}
		return r, nil
	case nil:
		return r, Nil
	default:
		return r, fmt.Errorf("unable to cast %#v of type %T to []any", val, val)
	}
}

func toMap(a any) (map[string]any, error) {
	r := make(map[string]any)
	switch m := a.(type) {
	case map[any]any:
		for k, val := range m {
			s, err := toString(k)
			if err != nil {
				return nil, err
			}
			r[s] = val
		}
		return r, nil
	case map[string]any:
		return m, nil
	case nil:
		return r, Nil
	default:
		return r, fmt.Errorf("unable to cast %#v of type %T to map[string]any", a, a)
	}
}

func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
