package gi2

import (
	"iter"
)

// Fold reduces the sequence of pairs to a single pair using the given function.
func Fold[V1, V2, R1, R2 any, F ~func(R1, R2, V1, V2) (R1, R2)](pairs iter.Seq2[V1, V2], i1 R1, i2 R2, accumulate F) (R1, R2) {
	r1, r2 := i1, i2

	for v1, v2 := range pairs {
		r1, r2 = accumulate(r1, r2, v1, v2)
	}

	return r1, r2
}

// FoldPack reduces the sequence of pairs to a single value using the given function.
func FoldPack[V1, V2, R any, F ~func(R, V1, V2) R](pairs iter.Seq2[V1, V2], initial R, accumulate F) R {
	result := initial

	for v1, v2 := range pairs {
		result = accumulate(result, v1, v2)
	}

	return result
}
