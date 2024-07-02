package statusx

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func NewCause(cause error) *Cause {
	causeProto, ok := cause.(proto.Message)
	if ok {
		causeAny, err := anypb.New(causeProto)
		if err != nil {
			panic(err)
		}
		return &Cause{Cause: &Cause_Error{Error: causeAny}}
	}
	return &Cause{Cause: &Cause_Message{Message: cause.Error()}}
}

func (x *Cause) Error() error {
	causeAny := x.GetError()
	if causeAny != nil {
		causeProto, err := causeAny.UnmarshalNew()
		if err != nil {
			panic(err)
		}
		if err, ok := causeProto.(error); ok {
			return err
		}
	}
	return errors.New(x.GetMessage())
}
