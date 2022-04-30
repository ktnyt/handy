package enums

// Filter returns the elements in the input slice that returns true when
// applied to the given filter function.
func Filter[E any](in []E, filter func(elem E) bool) []E {
	out := make([]E, 0, len(in)/2)
	for _, elem := range in {
		if filter(elem) {
			out = append(out, elem)
		}
	}
	return out
}
