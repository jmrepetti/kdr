package cherry

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Check2[T any](v T, err error) T {
	Check(err)
	return v
}
