package gi

import (
	"iter"
)

// Fold reduces the slice of values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
func Fold[V, R any, F ~func(R, V) R](values iter.Seq[V], initial R, accumulate F) R {
	result := initial

	for value := range values {
		result = accumulate(result, value)
	}

	return result
}
