# Box

Save Go structs to disk (or another place implementing StorageAdapter interface)


```go
adapter := NewFileStore("./box_storage/")
bx := NewBox(adapter)
// Register struct so it can be decoded later
bx.Register(JobTest{}) 
aJob := JobTest{ID: 2048, ShouldFail: true}
// Store return an ID
id, err := bx.Store(aJob)
// Use previous ID for fetching structs from storage
job, err := bx.Fetch(id)
assert.Equal(t, 2048, job.(JobTest).ID)
```

Box of interface T

```go

type Job interface {
    Perform() error
}

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

adapter := NewMapStore()
bx := NewBoxT[Job](adapter)
bx.Register(JobTest{})
aJob := JobTest{ID: 2048, ShouldFail: true}
id, err := bx.Store(aJob)
job, err := bx.Fetch(id)
assert.ErrorContains(t, job.Perform(), "Something went wrong")
```