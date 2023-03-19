package convx

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

func Int(a any) (int64, error) {
	a = reflectx.Indirect(a)
	ty := reflect.TypeOf(a)
	kind := ty.Kind()
	switch kind {
	case reflect.Int:
		i := a.(int)
		return int64(i), nil
	case reflect.Int8:
		i := a.(int8)
		return int64(i), nil
	case reflect.Int16:
		i := a.(int16)
		return int64(i), nil
	case reflect.Int32:
		i := a.(int32)
		return int64(i), nil
	case reflect.Int64:
		i := a.(int64)
		return i, nil
	case reflect.Uint:
		i := a.(uint)
		return int64(i), nil
	case reflect.Uint8:
		i := a.(uint8)
		return int64(i), nil
	case reflect.Uint16:
		i := a.(uint16)
		return int64(i), nil
	case reflect.Uint32:
		i := a.(uint32)
		return int64(i), nil
	case reflect.Uint64:
		i := a.(uint64)
		return int64(i), nil
	case reflect.Uintptr:
		p := a.(uintptr)
		return int64(p), nil
	case reflect.Float32:
		f := a.(float32)
		return int64(f), nil
	case reflect.Float64:
		f := a.(float64)
		return int64(f), nil
	case reflect.String:
		s := a.(string)
		return strconv.ParseInt(s, 10, 64)
	case reflect.Slice:
		elemKind := ty.Elem().Kind()
		if elemKind == reflect.Uint8 {
			s := a.([]byte)
			return strconv.ParseInt(string(s), 10, 64)
		}
	}
	return 0, fmt.Errorf("unable to convert %#v of type %T to int64", a, a)
}

//	Bool
//	Int
//	Int8
//	Int16
//	Int32
//	Int64
//	Uint
//	Uint8
//	Uint16
//	Uint32
//	Uint64
//	Uintptr
//	Float32
//	Float64
//	Array
//	Chan
//	Func
//	Interface
//	Map
//	Pointer
//	Slice
//	String
//	Struct
//	UnsafePointer
