package gi2

// IsSome returns true if the right value of the given trinity of values is true.
// Otherwise, it returns false.
// This function is useful for checking if a value is present in an optional pair of values.
func IsSome[V1, V2 any](_ V1, _ V2, some bool) bool {
	return some
}

// IsNone returns true if the right value of the given trinity of values is false.
// Otherwise, it returns false.
// This function is useful for checking if a value is absent in an optional pair of values.
func IsNone[V1, V2 any](_ V1, _ V2, some bool) bool {
	return !some
}
