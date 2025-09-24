package chimp

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var ops atomic.Int64

type JobTest struct {
	ID          int
	ShouldFail  bool
	ShouldPanic bool
}

func (j *JobTest) Perform() error {
	ops.Add(1)
	if j.ShouldPanic {
		panic("Something went panic")
	}
	if j.ShouldFail {
		return fmt.Errorf("Something went wrong")
	}
	return nil
}

func TestSchedule(t *testing.T) {
	sch := NewScheduler()
	job := JobTest{ID: 1}
	when := time.Now().Add(1000)
	assert.NoError(t, sch.Schedule(&job, when))
	assert.Equal(t, int64(1), sch.PendingJobsCount())
}

func TestScheduleError(t *testing.T) {
	sch := NewScheduler()
	job := JobTest{ID: 1}
	when := time.Now().Add(-1000)
	assert.Error(t, sch.Schedule(&job, when))
	assert.Equal(t, int64(0), sch.PendingJobsCount())
}

func TestJobError(t *testing.T) {
	ops.Store(0)
	sch := NewScheduler()
	sch.Start()
	job := JobTest{ID: 1, ShouldFail: true}
	assert.NoError(t, sch.ScheduleNow(&job))
	err := <-sch.Errors()
	assert.ErrorContains(t, err, "Something went wrong")
	assert.Equal(t, int64(1), sch.failedJobs.Load())
	sch.Stop()
}

func TestJobPanicRecover(t *testing.T) {
	ops.Store(0)
	sch := NewScheduler()
	sch.Start()
	job := JobTest{ID: 1, ShouldPanic: true}
	assert.NoError(t, sch.ScheduleNow(&job))
	err := <-sch.Errors()
	assert.ErrorContains(t, err, "Something went panic")
	assert.Equal(t, int64(1), sch.failedJobs.Load())
	sch.Stop()
}

func TestPendingJobsCount(t *testing.T) {
	ops.Store(0)
	sch := NewScheduler()
	jobsCount := 100
	sch.Start()
	for i := range jobsCount {
		assert.NoError(t, sch.ScheduleNow(&JobTest{ID: i}))
	}
	assert.Equal(t, int64(jobsCount), sch.PendingJobsCount())
	time.Sleep(2 * time.Second)
	assert.Equal(t, int64(0), sch.PendingJobsCount())
	assert.Equal(t, int64(jobsCount), ops.Load())
	sch.Stop()
}
