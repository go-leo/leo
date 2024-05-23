package cqrs

import "reflect"

var metadataType = reflect.TypeOf((*Metadata)(nil)).Elem()
