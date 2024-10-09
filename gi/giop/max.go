package giop

import (
	"golang.org/x/exp/constraints"

	"github.com/pamburus/go-mod/gi/gic"
)

// Max returns the maximum value of the left and right.
func Max[T constraints.Ordered](left, right T) T {
	return max(left, right)
}

// MaxBy returns a function that returns the maximum value of the left and right
// using provided key function for comparison.
func MaxBy[T any, K constraints.Ordered, Key ~func(T) K](key Key) func(T, T) T {
	return func(left, right T) T {
		if key(left) > key(right) {
			return left
		}

		return right
	}
}

// MaxByLess returns the maximum value of the left and right using Less method for comparison.
func MaxByLess[T gic.OrderedByLess[T]](left, right T) T {
	return MaxByLessFunc(T.Less)(left, right)
}

// MaxByLessFunc returns a function that returns the maximum value of the left and right
// using provided less function for comparison.
func MaxByLessFunc[T any, F ~func(T, T) bool](less F) func(left, right T) T {
	return func(left, right T) T {
		if less(right, left) {
			return left
		}

		return right
	}
}

// MaxByCompare returns the maximum value of the left and right using Compare method for comparison.
func MaxByCompare[T gic.OrderedByCompare[T]](left, right T) T {
	return MaxByCompareFunc(T.Compare)(left, right)
}

// MaxByCompareFunc returns a function that returns the maximum value of the left and right
// using provided compare function for comparison.
func MaxByCompareFunc[T any, F ~func(T, T) int](compare F) func(left, right T) T {
	return MaxByLessFunc(func(left, right T) bool {
		return compare(left, right) < 0
	})
}
