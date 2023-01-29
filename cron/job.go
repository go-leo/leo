package cron

import (
	"time"

	"github.com/robfig/cron/v3"
)

type Job struct {
	// 任务名
	name string
	// cron标准时间定义
	spec string
	// 定时执行的方法
	cmd func()
	// cron里一个Entry
	entryID cron.EntryID
	// Cron实例
	exec *cron.Cron
}

func NewJob(name, spec string, cmd cron.FuncJob) *Job {
	return &Job{name: name, spec: spec, cmd: cmd}
}

func (job *Job) Name() string {
	return job.name
}

func (job *Job) Spec() string {
	return job.spec
}

func (job *Job) EntryID() cron.EntryID {
	return job.entryID
}

// Stop 停止一个Job
func (job *Job) Stop() {
	if job.exec != nil {
		job.exec.Remove(job.EntryID())
	}
}

// Next time the job will run, or the zero time if Cron has not been
// started or this entry's schedule is unsatisfiable
func (job *Job) Next() time.Time {
	if job.exec != nil {
		return job.exec.Entry(job.EntryID()).Next
	}
	return time.Time{}
}

// Prev is the last time this job was run, or the zero time if never.
func (job *Job) Prev() time.Time {
	if job.exec != nil {
		return job.exec.Entry(job.EntryID()).Prev
	}
	return time.Time{}
}
