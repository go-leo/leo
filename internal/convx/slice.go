package convx

import (
	"fmt"
	"reflect"

	"github.com/go-leo/gox/reflectx"
)

func Slice(a any) ([]any, error) {
	a = reflectx.Indirect(a)
	val := reflect.ValueOf(a)
	kind := val.Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		r := make([]any, 0, val.Len())
		for i := 0; i < val.Len(); i++ {
			r = append(r, val.Index(i).Interface())
		}
		return r, nil
	}
	return nil, fmt.Errorf("unable to convert %#v of type %T to []any", a, a)
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
