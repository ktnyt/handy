package enums

// IndexFunc returns the index of the first element that matches the given
// check function.
func IndexFunc[Elem any](in []Elem, check func(elem Elem) bool) int {
	for i, elem := range in {
		if check(elem) {
			return i
		}
	}
	return -1
}

// Index returns the index of the first element that is equal to the given
// element.
func Index[Elem comparable](in []Elem, elem Elem) int {
	for i, v := range in {
		if v == elem {
			return i
		}
	}
	return -1
}

// Find returns the value of the first element that matches the given find
// function.
func Find[Elem any](in []Elem, find func(elem Elem) bool) Elem {
	return in[IndexFunc(in, find)]
}

// ContainsFunc reports whether an element that matches the given contains
// function is present in in.
func ContainsFunc[Elem any](in []Elem, contains func(elem Elem) bool) bool {
	return IndexFunc(in, contains) >= 0
}

// Contains reports whether an element v is present in in.
func Contains[Elem comparable](in []Elem, elem Elem) bool {
	return Index(in, elem) >= 0
}
