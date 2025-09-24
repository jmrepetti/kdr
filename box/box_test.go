package box

import (
	"fmt"
	"testing"

	"github.com/jmrepetti/kdr/chimp"
	"github.com/stretchr/testify/assert"
)

type JobTest struct {
	ID          int
	ShouldFail  bool
	ShouldPanic bool
}

func (j JobTest) Perform() error {
	if j.ShouldPanic {
		panic("Something went panic")
	}
	if j.ShouldFail {
		return fmt.Errorf("Something went wrong")
	}
	return nil
}

func TestBoxStore(t *testing.T) {
	adapter := NewMapStore()
	bx := NewBox(adapter)
	bx.Register(JobTest{})
	aJob := JobTest{ID: 2048, ShouldFail: true}
	id, err := bx.Store(aJob)
	assert.NoError(t, err)
	job, err := bx.Fetch(id)
	assert.NoError(t, err)
	assert.Equal(t, 2048, job.(JobTest).ID)
	assert.True(t, job.(JobTest).ShouldFail)
	assert.Error(t, PerformJon(job.(JobTest)))
}

func TestFileStore(t *testing.T) {
	adapter := NewFileStore("./box_storage/")
	bx := NewBox(adapter)
	bx.Register(JobTest{})
	aJob := JobTest{ID: 2048, ShouldFail: true}
	id, err := bx.Store(aJob)
	assert.NoError(t, err)
	job, err := bx.Fetch(id)
	assert.NoError(t, err)
	assert.Equal(t, 2048, job.(JobTest).ID)
	assert.ErrorContains(t, PerformJon(job.(JobTest)), "Something went wrong")
}

func PerformJon(j chimp.Job) error {
	return j.Perform()
}
