package schedule

import "context"

// Task is a task
type Task interface {
	// Name 任务名
	Name() string
	// Expression 任务时间规则表达式
	Expression() string
	// Invoke 执行任务
	Invoke(ctx context.Context)
}

type Invoker interface {
	Invoke(ctx context.Context)
}

type Middleware interface {
	Decorate(invoker Invoker) Invoker
}

type InvokerFunc func(ctx context.Context)

func (f InvokerFunc) Invoke(ctx context.Context) {
	f(ctx)
}

type MiddlewareFunc func(invoker Invoker) Invoker

func (f MiddlewareFunc) Decorate(invoker Invoker) Invoker {
	return f(invoker)
}

// Chain decorates the given Invoker with all middlewares in the chain.
func Chain(invoker Invoker, middlewares ...Middleware) Invoker {
	for i := len(middlewares) - 1; i >= 0; i-- {
		invoker = middlewares[i].Decorate(invoker)
	}
	return invoker
}
