package enums

// Reduce returns a value accumulated by repeatedly applying the given reduce
// function on each element of the input slice.
func Reduce[E any, A any](in []E, acc A, reduce func(acc A, elem E) A) A {
	for _, elem := range in {
		acc = reduce(acc, elem)
	}
	return acc
}

// MapReduce returns all of the values computed by Reduce, including the
// initial accumulator value.
func MapReduce[E any, A any](in []E, acc A, reduce func(acc A, elem E) A) []A {
	out := make([]A, len(in)+1)
	out[0] = acc
	for i, elem := range in {
		out[i+1] = reduce(out[i], elem)
	}
	return out
}
