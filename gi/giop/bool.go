package giop

// And returns true if both left and right are true.
func And(left, right bool) bool {
	return left && right
}

// Or returns true if either left or right is true.
func Or(left, right bool) bool {
	return left || right
}

// Xor returns true if either left or right is true, but not both.
func Xor(left, right bool) bool {
	return left != right
}
