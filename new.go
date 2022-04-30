package handy

// New returns a pointer initialized with the given value.
func New[T any](v T) *T {
	return &v
}
