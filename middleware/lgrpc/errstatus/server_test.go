package grpcerr

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestTimeoutStatus(t *testing.T) {
	o := ErrorFunc(func(err error) *status.Status {
		switch err {
		case nil:
			return nil
		case context.DeadlineExceeded:
			return status.New(codes.DeadlineExceeded, err.Error())
		default:
			return status.New(codes.Unknown, err.Error())
		}
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := UnaryServerInterceptor(o)(
		ctx,
		nil,
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			<-ctx.Done()
			return nil, context.DeadlineExceeded
		},
	)

	s, ok := status.FromError(err)
	if !ok {
		t.Errorf("Expected status error, got %+v", err)
	} else {
		if s.Code() != codes.DeadlineExceeded {
			t.Errorf("Expected status code %v, got %v", codes.DeadlineExceeded, s.Code())
		}
	}
}

func TestCancellationStatus(t *testing.T) {
	o := ErrorFunc(func(err error) *status.Status {
		switch err {
		case nil:
			return nil
		case context.Canceled:
			return status.New(codes.Canceled, err.Error())
		default:
			return status.New(codes.Unknown, err.Error())
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := UnaryServerInterceptor(o)(
		ctx,
		nil,
		&grpc.UnaryServerInfo{},
		func(ctx context.Context, req interface{}) (interface{}, error) {
			<-ctx.Done()
			return nil, context.Canceled
		},
	)

	s, ok := status.FromError(err)
	if !ok {
		t.Errorf("Expected status error, got %+v", err)
	} else {
		if s.Code() != codes.Canceled {
			t.Errorf("Expected status code %v, got %v", codes.Canceled, s.Code())
		}
	}
}
