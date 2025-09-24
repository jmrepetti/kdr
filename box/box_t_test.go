package box

import (
	"testing"

	"github.com/jmrepetti/kdr/chimp"
	"github.com/stretchr/testify/assert"
)

func TestBoxTStore(t *testing.T) {
	adapter := NewMapStore()
	bx := NewBoxT[chimp.Job](adapter)
	bx.Register(JobTest{})
	aJob := JobTest{ID: 2048, ShouldFail: true}
	id, err := bx.Store(aJob)
	assert.NoError(t, err)
	job, err := bx.Fetch(id)
	assert.NoError(t, err)
	assert.Equal(t, 2048, job.(JobTest).ID)
	assert.ErrorContains(t, job.Perform(), "Something went wrong")
}
