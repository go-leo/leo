package timex

import "time"

type Task struct {
	Name      string
	StartTime time.Time
	StopTime  time.Time
	Duration  time.Duration
}

type StopWatch struct {
	ID string
}

func (StopWatch) Start(taskName string) *Task {
	task := &Task{
		Name:      taskName,
		StartTime: time.Now(),
	}
	return task
}

func (StopWatch) Stop(task *Task) {
	task.StopTime = time.Now()
	task.Duration = task.StopTime.Sub(task.StartTime)
}
