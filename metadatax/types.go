package metadatax

import "reflect"

var Type = reflect.TypeOf((*Metadata)(nil)).Elem()
