package gi

import (
	"iter"

	"github.com/pamburus/go-mod/optional/optval"
)

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// Reduce returns [optval.None] if the sequence is empty.
func Reduce[V any, F ~func(V, V) V](values iter.Seq[V], accumulate F) optval.Value[V] {
	return ReduceWith(accumulate)(values)
}

// ReduceWith returns a function that reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// The function returned by ReduceWith returns [optval.None] if the sequence is empty.
func ReduceWith[V any, F ~func(V, V) V](accumulate F) func(values iter.Seq[V]) optval.Value[V] {
	return func(values iter.Seq[V]) optval.Value[V] {
		return Fold(values, optval.None[V](), func(r optval.Value[V], v V) optval.Value[V] {
			if rv, ok := r.Unwrap(); ok {
				return optval.Some(accumulate(rv, v))
			}

			return optval.Some(v)
		})
	}
}
