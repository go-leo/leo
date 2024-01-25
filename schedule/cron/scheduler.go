package cron

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/robfig/cron/v3"
	"golang.org/x/exp/maps"

	"codeup.aliyun.com/qimao/leo/leo/actuator/health"

	"codeup.aliyun.com/qimao/leo/leo/actuator"

	"codeup.aliyun.com/qimao/leo/leo/schedule"
)

var _ schedule.Scheduler = new(Scheduler)

// Scheduler is task scheduler
type Scheduler struct {
	cron        *cron.Cron
	taskJobMap  map[schedule.Task]*cronJob
	middlewares []schedule.Middleware
	mutex       sync.RWMutex
	run         atomic.Bool
	alive       atomic.Bool
}

func NewScheduler(opts ...Option) *Scheduler {
	o := new(options)
	o.apply(opts...)
	o.init()
	scheduler := &Scheduler{
		cron:        cron.New(o.cronOptions()...),
		taskJobMap:  make(map[schedule.Task]*cronJob),
		middlewares: o.Middlewares,
	}
	err := scheduler.AddTasks(o.Task...)
	if err != nil {
		panic(fmt.Errorf("failed to add task, %w", err))
	}
	return scheduler
}

func (s *Scheduler) Run(ctx context.Context) error {
	if !s.run.CompareAndSwap(false, true) {
		return errors.New("scheduler is running")
	}
	s.alive.Store(true)
	defer s.alive.Store(false)
	s.mutex.RLock()
	// 异步开启cron
	s.startCron(ctx)
	s.mutex.RUnlock()
	// 等待context被cancel
	<-ctx.Done()
	// 停止cron
	s.stopCron(ctx)
	return nil
}

func (s *Scheduler) AddTasks(tasks ...schedule.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	var errs []error
	for _, task := range tasks {
		if _, ok := s.taskJobMap[task]; ok {
			return errors.New("task already exists")
		}
		job := &cronJob{
			Task:        task,
			Middlewares: s.middlewares,
		}
		err := s.addJob(job)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		s.taskJobMap[task] = job
	}

	return errors.Join(errs...)
}

func (s *Scheduler) RemoveTasks(tasks ...schedule.Task) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, task := range tasks {
		job, ok := s.taskJobMap[task]
		if !ok {
			continue
		}
		s.removeJob(job)
		delete(s.taskJobMap, task)
	}
	return nil
}

func (s *Scheduler) Tasks() []schedule.Task {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return maps.Keys(s.taskJobMap)
}

func (s *Scheduler) ActuatorHandler() actuator.Handler {
	return &actuatorHandler{scheduler: s}
}

func (s *Scheduler) HealthChecker() health.Checker {
	return &healthChecker{scheduler: s}
}

func (s *Scheduler) startCron(ctx context.Context) {
	// 异步运行cron
	s.cron.Start()
	runtime.Gosched()
}

func (s *Scheduler) stopCron(ctx context.Context) {
	// 移除所有job
	for _, job := range s.taskJobMap {
		s.removeJob(job)
	}
	// 停止cron
	<-s.cron.Stop().Done()
}

func (s *Scheduler) addJob(job *cronJob) error {
	entryID, err := s.cron.AddJob(job.Task.Expression(), job)
	if err != nil {
		return err
	}
	job.EntryID = entryID
	return nil
}

func (s *Scheduler) removeJob(job *cronJob) {
	s.cron.Remove(job.EntryID)
}

func (s *Scheduler) isAlive() bool {
	return s.run.Load()
}
