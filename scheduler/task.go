package scheduler

import "context"

// Task 任务
type Task interface {
	// Name 任务名
	Name() string
	// Schedule 任务时间规则
	Schedule() string
	// Invoke 执行任务
	Invoke(ctx context.Context)
}
