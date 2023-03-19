package convx

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

func String(a any) (string, error) {
	stringer, ok := a.(fmt.Stringer)
	if ok {
		return stringer.String(), nil
	}

	a = reflectx.Indirect(a)
	ty := reflect.TypeOf(a)
	kind := ty.Kind()
	switch kind {
	case reflect.Bool:
		b := a.(bool)
		return strconv.FormatBool(b), nil
	case reflect.Int:
		i := a.(int)
		return strconv.FormatInt(int64(i), 10), nil
	case reflect.Int8:
		i := a.(int8)
		return strconv.FormatInt(int64(i), 10), nil
	case reflect.Int16:
		i := a.(int16)
		return strconv.FormatInt(int64(i), 10), nil
	case reflect.Int32:
		i := a.(int32)
		return strconv.FormatInt(int64(i), 10), nil
	case reflect.Int64:
		i := a.(int64)
		return strconv.FormatInt(i, 10), nil
	case reflect.Uint:
		i := a.(uint)
		return strconv.FormatUint(uint64(i), 10), nil
	case reflect.Uint8:
		i := a.(uint8)
		return strconv.FormatUint(uint64(i), 10), nil
	case reflect.Uint16:
		i := a.(uint16)
		return strconv.FormatUint(uint64(i), 10), nil
	case reflect.Uint32:
		i := a.(uint32)
		return strconv.FormatUint(uint64(i), 10), nil
	case reflect.Uint64:
		i := a.(uint64)
		return strconv.FormatUint(i, 10), nil
	case reflect.Uintptr:
		p := a.(uintptr)
		return strconv.FormatUint(uint64(p), 10), nil
	case reflect.Float32:
		f := a.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 64), nil
	case reflect.Float64:
		f := a.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64), nil
	case reflect.String:
		s := a.(string)
		return s, nil
	case reflect.Slice:
		elemKind := ty.Elem().Kind()
		if elemKind == reflect.Uint8 {
			s := a.([]byte)
			return string(s), nil
		}
	}
	return "", fmt.Errorf("unable to convert %#v of type %T to string", a, a)
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
