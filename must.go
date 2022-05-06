package handy

// Must panics if err is not nil. Otherwise val is returned.
func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}
