// Package cmpx provides types and functions related to comparing values and beyond.
package cmpx

// IfElse returns t if condition is true, otherwise it returns f.
func IfElse[T any](condition bool, t, f T) T {
	if condition {
		return t
	}

	return f
}
