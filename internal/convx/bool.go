package convx

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/go-leo/gox/reflectx"
)

func Bool(a any) (bool, error) {
	a = reflectx.Indirect(a)
	ty := reflect.TypeOf(a)
	kind := ty.Kind()
	switch kind {
	case reflect.Bool:
		b := a.(bool)
		return b, nil
	case reflect.String:
		s := a.(string)
		return strconv.ParseBool(s)
	case reflect.Slice:
		elemKind := ty.Elem().Kind()
		if elemKind == reflect.Uint8 {
			s := a.([]byte)
			return strconv.ParseBool(string(s))
		}
	}
	return false, fmt.Errorf("unable to convert %#v of type %T to bool", a, a)
}
