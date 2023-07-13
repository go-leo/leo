package stream

import "context"

// Invoker is called by Interceptor.
type Invoker func(ctx context.Context, msg *Message, channel Channel) error

// Interceptor intercepts the message.
type Interceptor func(ctx context.Context, msg *Message, channel Channel, invoker Invoker) error

// chainInterceptors chains all unary client interceptors into one.
func chainInterceptors(interceptors []Interceptor) Interceptor {
	var chainedInt Interceptor
	if len(interceptors) == 0 {
		chainedInt = nil
	} else if len(interceptors) == 1 {
		chainedInt = interceptors[0]
	} else {
		chainedInt = func(ctx context.Context, msg *Message, channel Channel, invoker Invoker) error {
			return interceptors[0](ctx, msg, channel, getInvoker(interceptors, 0, invoker))
		}
	}
	return chainedInt
}

// getInvoker recursively generate the chained invoker.
func getInvoker(interceptors []Interceptor, curr int, finalInvoker Invoker) Invoker {
	if curr == len(interceptors)-1 {
		return finalInvoker
	}
	return func(ctx context.Context, msg *Message, channel Channel) error {
		return interceptors[curr+1](ctx, msg, channel, getInvoker(interceptors, curr+1, finalInvoker))
	}
}
