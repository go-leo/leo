package grpcerr

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type options struct {
	errorFunc func(err error) *status.Status
}

func (o *options) apply(opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

type Option func(o *options)

func defaultOptions() *options {
	return &options{
		errorFunc: func(err error) *status.Status {
			switch err {
			case nil:
				return nil
			case context.DeadlineExceeded:
				return status.New(codes.DeadlineExceeded, err.Error())
			case context.Canceled:
				return status.New(codes.Canceled, err.Error())
			default:
				if se, ok := err.(interface{ GRPCStatus() *status.Status }); ok {
					return se.GRPCStatus()
				} else {
					return status.New(codes.Unknown, err.Error())
				}
			}
		},
	}
}

func ErrorFunc(errorFunc func(err error) *status.Status) Option {
	return func(o *options) {
		o.errorFunc = errorFunc
	}
}
