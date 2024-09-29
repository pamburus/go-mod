package gi

import "iter"

// PairFold returns an iterator adapter over pairs that uses a transform function to transform pairs to values.
func PairFold[V1, V2, R any, F ~func(V1, V2) R](pairs iter.Seq2[V1, V2], fold F) iter.Seq[R] {
	return func(yield func(R) bool) {
		for v1, v2 := range pairs {
			if !yield(fold(v1, v2)) {
				return
			}
		}
	}
}
