package gi

import (
	"iter"

	"github.com/pamburus/go-mod/optional"
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

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// Reduce returns [optional.None] if the sequence is empty.
func Reduce[V any, F ~func(V, V) V](values iter.Seq[V], accumulate F) optional.Value[V] {
	return Fold(values, optional.None[V](), func(r optional.Value[V], v V) optional.Value[V] {
		if rv, ok := r.Unwrap(); ok {
			return optional.Some(accumulate(rv, v))
		}

		return optional.Some(v)
	})
}
