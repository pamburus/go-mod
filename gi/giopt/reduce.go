package giopt

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
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
func ReduceWith[V any, F ~func(V, V) V](accumulate F) func(iter.Seq[V]) optval.Value[V] {
	return func(values iter.Seq[V]) optval.Value[V] {
		return optval.New(gi.Reduce(values, accumulate))
	}
}
