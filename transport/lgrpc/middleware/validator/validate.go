package validator

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(req interface{}) error {
	switch v := req.(type) {
	case interface{ Validate() error }:
		if err := v.Validate(); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
	case interface{ Validate(all bool) error }:
		if err := v.Validate(false); err != nil {
			return status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return nil
}
