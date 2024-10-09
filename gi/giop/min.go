package giop

import (
	"golang.org/x/exp/constraints"

	"github.com/pamburus/go-mod/gi/gic"
)

// Min returns the minimum value of the left and right.
func Min[T constraints.Ordered](left, right T) T {
	return min(left, right)
}

// MinBy returns a function that returns the minimum value of the left and right
// using provided key function for comparison.
func MinBy[T any, K constraints.Ordered, Key ~func(T) K](key Key) func(T, T) T {
	return func(left, right T) T {
		if key(left) < key(right) {
			return left
		}

		return right
	}
}

// MinByLess returns the minimum value of the left and right using Less method for comparison.
func MinByLess[T gic.OrderedByLess[T]](left, right T) T {
	return MinByLessFunc(T.Less)(left, right)
}

// MinByLessFunc returns a function that returns the minimum value of the left and right
// using provided less function for comparison.
func MinByLessFunc[T any, F ~func(T, T) bool](less F) func(left, right T) T {
	return func(left, right T) T {
		if less(left, right) {
			return left
		}

		return right
	}
}

// MinByCompare returns the minimum value of the left and right using Compare method for comparison.
func MinByCompare[T gic.OrderedByCompare[T]](left, right T) T {
	return MinByCompareFunc(T.Compare)(left, right)
}

// MinByCompareFunc returns a function that returns the minimum value of the left and right
// using provided compare function for comparison.
func MinByCompareFunc[T any, F ~func(T, T) int](compare F) func(left, right T) T {
	return MinByLessFunc(func(left, right T) bool {
		return compare(left, right) < 0
	})
}
