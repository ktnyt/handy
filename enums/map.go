package enums

// Map returns a slice with each of the element in the input slice mapped with
// the given convert function.
func Map[In any, Out any](in []In, convert func(elem In) Out) []Out {
	out := make([]Out, len(in))
	for i, elem := range in {
		out[i] = convert(elem)
	}
	return out
}
