package chimp

import (
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"
)

type Job interface {
	Perform() error
}

type Scheduler struct {
	q             map[int64][]Job
	qMtx          *sync.RWMutex
	errChan       chan error
	stopChan      chan struct{}
	pendingJobs   atomic.Int64
	processedJobs atomic.Int64
	failedJobs    atomic.Int64
	ticker        *time.Ticker
}

func (sch *Scheduler) recover() {
	if r := recover(); r != nil {
		sch.failedJobs.Add(1)
		sch.errChan <- fmt.Errorf("recovered from panic: %v", r)
	}
}

// TODO: accept options
func NewScheduler() *Scheduler {
	return &Scheduler{
		q:        make(map[int64][]Job),
		qMtx:     &sync.RWMutex{},
		errChan:  make(chan error),
		stopChan: make(chan struct{}),
		ticker:   time.NewTicker(500 * time.Millisecond),
	}
}

func (sch *Scheduler) Errors() <-chan error {
	return sch.errChan
}

func (sch *Scheduler) Start() {
	go sch.worker()
}

func (sch *Scheduler) PerformJob(job Job) {
	defer sch.recover()
	defer sch.processedJobs.Add(1)
	defer sch.pendingJobs.Add(-1)
	if err := job.Perform(); err != nil {
		sch.failedJobs.Add(1)
		sch.errChan <- err
	}
}

func (sch *Scheduler) worker() {
	for {
		select {
		case <-sch.stopChan:
			return
		case t := <-sch.ticker.C:
			tunix := t.Unix()
			jobs := sch.getJobsAt(tunix)
			for _, job := range jobs {
				go sch.PerformJob(job)
			}
		}
	}
}

func (sch *Scheduler) Stop() {
	slog.Info(fmt.Sprintf("Stopping..."))
	close(sch.stopChan)
	close(sch.errChan)
	slog.Info(fmt.Sprintf("= Pending Jobs: %d", sch.PendingJobsCount()))
	slog.Info(fmt.Sprintf("= Processed Jobs: %d", sch.processedJobs.Load()))
	slog.Info(fmt.Sprintf("= Failed Jobs: %d", sch.failedJobs.Load()))
}

func (sch *Scheduler) ScheduleNow(job Job) error {
	return sch.Schedule(job, time.Now().Add(1*time.Second))
}

func (sch *Scheduler) addJobAt(job Job, when int64) error {
	sch.qMtx.Lock()
	defer sch.qMtx.Unlock()
	sch.pendingJobs.Add(1)
	sch.q[when] = append(sch.q[when], job)
	return nil
}

func (sch *Scheduler) getJobsAt(when int64) []Job {
	sch.qMtx.Lock()
	defer sch.qMtx.Unlock()
	jobs, ok := sch.q[when]
	if ok {
		delete(sch.q, when)
		return jobs
	} else {
		slog.Debug(fmt.Sprintf("No job at %v", when))
		return []Job{}
	}
}

func (sch *Scheduler) Schedule(job Job, when time.Time) error {
	now := time.Now()
	if when.Before(now) {
		return fmt.Errorf("schedule in the past. now: %v, when: %v", now.Unix(), when.Unix())
	}
	return sch.addJobAt(job, when.Unix())
}

func (sch *Scheduler) PendingJobsCount() int64 {
	return sch.pendingJobs.Load()
}
