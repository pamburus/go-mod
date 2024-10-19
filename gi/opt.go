package gi

// IsSome returns true if the right value of the given pair of values is true.
// Otherwise, it returns false.
// This function is useful for checking if a value is present in an optional value.
func IsSome[V any](_ V, some bool) bool {
	return some
}

// IsNone returns true if the right value of the given pair of values is false.
// Otherwise, it returns false.
// This function is useful for checking if a value is absent in an optional value.
func IsNone[V any](_ V, some bool) bool {
	return !some
}
