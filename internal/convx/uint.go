package convx

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

func Uint(a any) (uint64, error) {
	a = reflectx.Indirect(a)
	ty := reflect.TypeOf(a)
	kind := ty.Kind()
	switch kind {
	case reflect.Int:
		i := a.(int)
		return uint64(i), nil
	case reflect.Int8:
		i := a.(int8)
		return uint64(i), nil
	case reflect.Int16:
		i := a.(int16)
		return uint64(i), nil
	case reflect.Int32:
		i := a.(int32)
		return uint64(i), nil
	case reflect.Int64:
		i := a.(int64)
		return uint64(i), nil
	case reflect.Uint:
		i := a.(uint)
		return uint64(i), nil
	case reflect.Uint8:
		i := a.(uint8)
		return uint64(i), nil
	case reflect.Uint16:
		i := a.(uint16)
		return uint64(i), nil
	case reflect.Uint32:
		i := a.(uint32)
		return uint64(i), nil
	case reflect.Uint64:
		i := a.(uint64)
		return i, nil
	case reflect.Uintptr:
		p := a.(uintptr)
		return uint64(p), nil
	case reflect.Float32:
		f := a.(float32)
		return uint64(f), nil
	case reflect.Float64:
		f := a.(float64)
		return uint64(f), nil
	case reflect.String:
		s := a.(string)
		return strconv.ParseUint(s, 10, 64)
	case reflect.Slice:
		elemKind := ty.Elem().Kind()
		if elemKind == reflect.Uint8 {
			s := a.([]byte)
			return strconv.ParseUint(string(s), 10, 64)
		}
	}
	return 0, fmt.Errorf("unable to convert %#v of type %T to uint64", a, a)
}
