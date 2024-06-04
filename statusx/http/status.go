package http

import (
	"github.com/go-leo/gox/errorx"
	"google.golang.org/protobuf/types/known/anypb"
)

var AnyProto *anypb.Any

func init() {
	AnyProto = errorx.Ignore(anypb.New(&Status{}))
}
