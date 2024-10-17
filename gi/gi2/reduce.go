package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/optional/optpair"
)

// Reduce reduces the values to a value which is the accumulated result of running accumulate function
// on each element where each successive invocation is supplied the return value of the previous one.
//
// Reduce returns [optval.None] if the sequence is empty.
func Reduce[V1, V2 any, F ~func(V1, V2, V1, V2) (V1, V2)](values iter.Seq2[V1, V2], accumulate F) optpair.Pair[V1, V2] {
	return Fold1(values, optpair.None[V1, V2](), func(r optpair.Pair[V1, V2], v1 V1, v2 V2) optpair.Pair[V1, V2] {
		if rv1, rv2, ok := r.Unwrap(); ok {
			return optpair.Some(accumulate(rv1, rv2, v1, v2))
		}

		return optpair.Some(v1, v2)
	})
}
