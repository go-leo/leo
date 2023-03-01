package schedule

import "context"

// Task is a task
type Task interface {
	// Name 任务名
	Name() string
	// ID 任务ID
	ID() string
	// Expression 任务时间规则表达式
	Expression() string
	// Invoke 执行任务
	Invoke(ctx context.Context)
}
