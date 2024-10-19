package gi

import (
	"iter"
)

// Fold reduces the slice of values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
func Fold[V, R any, F ~func(R, V) R](values iter.Seq[V], initial R, accumulate F) R {
	return FoldWith(initial, accumulate)(values)
}

// FoldWith returns a function that can be used to fold sequences of values.
func FoldWith[V, R any, F ~func(R, V) R](initial R, accumulate F) func(iter.Seq[V]) R {
	return func(values iter.Seq[V]) R {
		result := initial

		for value := range values {
			result = accumulate(result, value)
		}

		return result
	}
}
