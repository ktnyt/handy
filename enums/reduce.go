package enums

// Reduce returns a value accumulated by repeatedly applying the given reduce
// function on each element of the input slice.
func Reduce[Elem any, Acc any](in []Elem, acc Acc, reduce func(acc Acc, elem Elem) Acc) Acc {
	for _, elem := range in {
		acc = reduce(acc, elem)
	}
	return acc
}

// MapReduce returns all of the values computed by Reduce, including the
// initial accumulator value.
func MapReduce[Elem any, Acc any](in []Elem, acc Acc, reduce func(acc Acc, elem Elem) Acc) []Acc {
	out := make([]Acc, len(in)+1)
	out[0] = acc
	for i, elem := range in {
		out[i+1] = reduce(out[i], elem)
	}
	return out
}
