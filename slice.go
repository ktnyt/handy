package handy

// Slice returns a slice with the given elements.
func Slice[T any](elements ...T) []T {
	return elements
}
