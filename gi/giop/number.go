package giop

import (
	"math"

	"golang.org/x/exp/constraints"

	"github.com/pamburus/go-mod/gi/gic"
)

// Add returns sum of left and right.
func Add[T gic.Number](left T, right T) T {
	return left + right
}

// Subtract returns difference of left and right.
func Subtract[T gic.Number](left T, right T) T {
	return left - right
}

// Multiply returns product of left and right.
func Multiply[T gic.Number](left T, right T) T {
	return left * right
}

// Divide returns integer part of left / right.
func Divide[T gic.Number](left T, right T) T {
	return left / right
}

// IntMod returns integer remainder of left / right.
func IntMod[T constraints.Integer](left T, right T) T {
	return left % right
}

// Mod returns math.Mod(left, right).
func Mod[T constraints.Float](left T, right T) T {
	return T(math.Mod(float64(left), float64(right)))
}

// BinaryAnd returns the binary "and" of the left and right.
func BinaryAnd[T constraints.Integer](left, right T) T {
	return left & right
}

// BinaryOr returns the binary "or" of the left and right.
func BinaryOr[T constraints.Integer](left, right T) T {
	return left | right
}

// BinaryXor returns the binary "xor" of the left and right.
func BinaryXor[T constraints.Integer](left, right T) T {
	return left ^ right
}
