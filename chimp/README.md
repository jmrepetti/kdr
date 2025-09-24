# Chimp

A simple, map-based job scheduler that lets you schedule tasks down to the second.

USAGE:

```go
package main

import (
	"fmt"
	"log/slog"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/jmrepetti/kdr/antenna"
	"github.com/jmrepetti/kdr/chimp"
)

var ops atomic.Int64

type JobTest struct {
	ID int
}

func (j *JobTest) Perform() error {
	ops.Add(1)
	fmt.Printf("DONE JOB %d: %v\n", j.ID, time.Now())
	return nil
}

func main() {
	ant := antenna.NewAntenna(syscall.SIGINT, syscall.SIGTERM)

	s := chimp.NewScheduler()
	s.Start()

    aJob := &JobTest{ID: i}
    when := time.Now().Add(10*time.Second)
    
    err := s.Schedule(aJob, when)

    if err != nil {
        slog.Error(err.Error())
    }

	go func() {
		for err := range s.Errors() {
			slog.Error(fmt.Sprintf("Job error: %v\n", err))
		}
	}()

	<-ant.Wait()
    s.Stop()
}
```