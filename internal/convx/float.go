package convx

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

func Float(a any) (float64, error) {
	a = reflectx.Indirect(a)
	ty := reflect.TypeOf(a)
	kind := ty.Kind()
	switch kind {
	case reflect.Int:
		i := a.(int)
		return float64(i), nil
	case reflect.Int8:
		i := a.(int8)
		return float64(i), nil
	case reflect.Int16:
		i := a.(int16)
		return float64(i), nil
	case reflect.Int32:
		i := a.(int32)
		return float64(i), nil
	case reflect.Int64:
		i := a.(int64)
		return float64(i), nil
	case reflect.Uint:
		i := a.(uint)
		return float64(i), nil
	case reflect.Uint8:
		i := a.(uint8)
		return float64(i), nil
	case reflect.Uint16:
		i := a.(uint16)
		return float64(i), nil
	case reflect.Uint32:
		i := a.(uint32)
		return float64(i), nil
	case reflect.Uint64:
		i := a.(uint64)
		return float64(i), nil
	case reflect.Uintptr:
		p := a.(uintptr)
		return float64(p), nil
	case reflect.Float32:
		f := a.(float32)
		return float64(f), nil
	case reflect.Float64:
		f := a.(float64)
		return f, nil
	case reflect.String:
		s := a.(string)
		return strconv.ParseFloat(s, 64)
	case reflect.Slice:
		elemKind := ty.Elem().Kind()
		if elemKind == reflect.Uint8 {
			s := a.([]byte)
			return strconv.ParseFloat(string(s), 64)
		}
	}
	return 0, fmt.Errorf("unable to convert %#v of type %T to float64", a, a)
}
