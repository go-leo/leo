package reflectx

import (
	"reflect"
)

func BaseKind(a any) reflect.Kind {
	a = Indirect(a)
	return reflect.TypeOf(a).Kind()
}
