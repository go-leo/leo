package recovery

import (
	"context"
	"fmt"

	"codeup.aliyun.com/qimao/leo/leo/stream"

	"codeup.aliyun.com/qimao/leo/leo/internal/gox/runtimex"
)

func Interceptor(handles ...func(context.Context, any) error) stream.Interceptor {
	var handle func(context.Context, any) error
	if len(handles) == 0 {
		handle = func(ctx context.Context, p any) error {
			return fmt.Errorf("panic triggered: %+v, stack: %s", p, runtimex.Stack(0))
		}
	} else {
		handle = handles[0]
	}

	return func(ctx context.Context, msg *stream.Message, channel stream.Channel, invoker stream.Invoker) (err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				err = handle(ctx, r)
			}
		}()

		err = invoker(ctx, msg, channel)
		panicked = false
		return err
	}
}
