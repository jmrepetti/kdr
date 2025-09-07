# Cherry package


Check panics if the provided error is not nil.
Useful for failing fast on unexpected errors

```go
Check(errors.New('something went wrong)) //=> panics
```

Check2 returns the value if there is no error, otherwise panics.

It is a convenience wrapper around `Check` for functions that return (T, error).

Example:

```go
val := Check2(someFuncThatReturnsValueAndError())
```
