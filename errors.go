package handy

import "errors"

// As is an inline form of errors.As.
func As[Error error](err error) (Error, bool) {
	var as Error
	ok := errors.As(err, &as)
	return as, ok
}

// IsType reports whether or not the type of any error in err's chain matches
// the Error type.
func IsType[Error error](err error) bool {
	_, ok := As[Error](err)
	return ok
}
