package enums

// Map returns a slice with each of the element in the input slice mapped with
// the given convert function.
func Map[E1 any, E2 any](in []E1, convert func(elem E1) E2) []E2 {
	out := make([]E2, len(in))
	for i, elem := range in {
		out[i] = convert(elem)
	}
	return out
}
