package cherry

// Check panics if the provided error is not nil.
// Useful for failing fast on unexpected errors
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// Check2 returns the value `v` if there is no error, otherwise panics.
// It is a convenience wrapper around `Check` for functions that return (T, error).
// Example: `val := Check2(someFuncThatReturnsValueAndError())`
func Check2[T any](v T, err error) T {
	Check(err)
	return v
}
