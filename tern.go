package handy

// Tern implements the ternary operator. It will return truthy if the condition
// is true and falsey if the condition is false.
func Tern[T any](cond bool, truthy, falsey T) T {
	if cond {
		return truthy
	}
	return falsey
}
