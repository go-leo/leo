package convx

import (
	"fmt"
	"reflect"

	"github.com/go-leo/gox/reflectx"
)

func Map(a any) (map[any]any, error) {
	a = reflectx.Indirect(a)
	val := reflect.ValueOf(a)
	kind := val.Kind()
	switch kind {
	case reflect.Map:
		r := make(map[any]any, val.Len())
		mapIter := val.MapRange()
		for mapIter.Next() {
			key := mapIter.Key()
			val := mapIter.Value()
			r[key.Interface()] = val.Interface()
		}
		return r, nil
	}
	return nil, fmt.Errorf("unable to convert %#v of type %T to map[any]any", a, a)
}
