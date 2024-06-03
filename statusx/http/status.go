package http

import (
	"github.com/go-leo/gox/errorx"
	"google.golang.org/protobuf/types/known/anypb"
)

var AnyProto = errorx.Ignore(anypb.New(&Status{}))
