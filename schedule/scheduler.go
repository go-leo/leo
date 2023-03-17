package schedule

import (
	"context"
)

// Scheduler is task scheduler
type Scheduler interface {
	// Run 开始运行
	Run(ctx context.Context) error
	// AddTasks 添加task
	AddTasks(tasks ...Task) error
	// RemoveTasks 删除task
	RemoveTasks(tasks ...Task) error
	// Tasks 任务列表
	Tasks() []Task
}
