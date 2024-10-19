package gi2opt

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// Reduce returns [optval.None] if the sequence is empty.
func Reduce[V1, V2 any, F ~func(V1, V2, V1, V2) (V1, V2)](values iter.Seq2[V1, V2], accumulate F) optpair.Pair[V1, V2] {
	return ReduceWith(accumulate)(values)
}

// ReduceWith returns a function that reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
func ReduceWith[V1, V2 any, F ~func(V1, V2, V1, V2) (V1, V2)](accumulate F) func(iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return func(values iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
		return optpair.New(gi2.Reduce(values, accumulate))
	}
}
