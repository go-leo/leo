package reflectx_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/go-leo/gox/reflectx"
)

func TestIndirect(t *testing.T) {
	var a = 10
	var aptr = &a
	fmt.Printf("%T %T \n", aptr, reflectx.Indirect(aptr))
	var aaptr = &aptr
	fmt.Printf("%T %T \n", aaptr, reflectx.Indirect(aaptr))
}

func TestIndirectAlias(t *testing.T) {
	type A int
	var a A = 10
	var aptr = &a
	fmt.Printf("%T %T \n", aptr, reflectx.Indirect(aptr))
	var aaptr = &aptr
	fmt.Printf("%T %T \n", aaptr, reflectx.Indirect(aaptr))
	var aaaptr = &aaptr
	fmt.Printf("%T %T \n", aaaptr, reflectx.Indirect(aaaptr))
	ty := reflect.TypeOf(a)
	fmt.Println(ty)
	fmt.Println(ty.Kind())

	vl := reflect.ValueOf(a)
	fmt.Println(vl)
	fmt.Println(vl.Kind())
	fmt.Println(vl.CanFloat())
	fmt.Println(vl.CanInt())
	fmt.Println(vl.CanUint())
	fmt.Println(vl.CanInterface())

}

func A(arg any) {
	k := v.kind()
	p := v.ptr
	switch k {
	case Int:
		return int64(*(*int)(p))
	case Int8:
		return int64(*(*int8)(p))
	case Int16:
		return int64(*(*int16)(p))
	case Int32:
		return int64(*(*int32)(p))
	case Int64:
		return *(*int64)(p)
	}
}

func TestIndirectAliasPtr(t *testing.T) {
	type A *int
	var n = 10
	var a A = &n
	var aptr = &a
	fmt.Printf("%T %T \n", aptr, reflectx.Indirect(aptr))
	var aaptr = &aptr
	fmt.Printf("%T %T \n", aaptr, reflectx.Indirect(aaptr))
}
