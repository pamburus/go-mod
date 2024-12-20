package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
)

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// Reduce returns [optval.None] if the sequence is empty.
func Reduce[V1, V2 any, F ~func(V1, V2, V1, V2) (V1, V2)](values iter.Seq2[V1, V2], accumulate F) (V1, V2, bool) {
	return ReduceWith(accumulate)(values)
}

// ReduceWith returns a function that reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
func ReduceWith[V1, V2 any, F ~func(V1, V2, V1, V2) (V1, V2)](accumulate F) func(iter.Seq2[V1, V2]) (V1, V2, bool) {
	return func(values iter.Seq2[V1, V2]) (V1, V2, bool) {
		type pair struct {
			v1 V1
			v2 V2
		}

		pack := func(v1 V1, v2 V2) pair {
			return pair{v1, v2}
		}

		result, ok := gi.Reduce(Pack(values, pack), func(l, r pair) pair {
			return pack(accumulate(l.v1, l.v2, r.v1, r.v2))
		})

		return result.v1, result.v2, ok
	}
}
