package scheduler

import (
	"context"
)

// Scheduler is task scheduler
type Scheduler interface {
	// Start 开始运行
	Start(ctx context.Context) error
	// Stop 停止运行
	Stop(ctx context.Context) error
	// AppendTask 追加task
	AppendTask(task Task) Scheduler
	// RemoveTask 删除task
	RemoveTask(task Task) Scheduler
	// Tasks 任务列表
	Tasks() []Task
}
