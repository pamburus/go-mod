package gi2

import "iter"

// PairFold returns an iterator that yields values by applying the provided fold function to pairs.
func PairFold[V1, V2, R any, F ~func(V1, V2) R](pairs iter.Seq2[V1, V2], fold F) iter.Seq[R] {
	return PairFoldWith(fold)(pairs)
}

// PairFoldWith returns a function that transforms an iterator sequence over pairs to an iterator sequence over values.
// It uses the provided fold function to transform pairs to values.
func PairFoldWith[V1, V2, R any, F ~func(V1, V2) R](fold F) func(iter.Seq2[V1, V2]) iter.Seq[R] {
	return func(pairs iter.Seq2[V1, V2]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v1, v2 := range pairs {
				if !yield(fold(v1, v2)) {
					return
				}
			}
		}
	}
}

// PairUnfold returns an iterator that yields pairs by applying the provided unfold function to values.
func PairUnfold[V, R1, R2 any, F ~func(V) (R1, R2)](values iter.Seq[V], unfold F) iter.Seq2[R1, R2] {
	return PairUnfoldWith(unfold)(values)
}

// PairUnfoldWith returns a function that transforms an iterator sequence over values to an iterator sequence over pairs.
// It uses the provided unfold function to transform values to pairs.
func PairUnfoldWith[V, R1, R2 any, F ~func(V) (R1, R2)](unfold F) func(iter.Seq[V]) iter.Seq2[R1, R2] {
	return func(values iter.Seq[V]) iter.Seq2[R1, R2] {
		return func(yield func(R1, R2) bool) {
			for v := range values {
				if !yield(unfold(v)) {
					return
				}
			}
		}
	}
}

// PairSwap returns an iterator that yields pairs with swapped values.
func PairSwap[V1, V2 any](pairs iter.Seq2[V1, V2]) iter.Seq2[V2, V1] {
	return func(yield func(V2, V1) bool) {
		for v1, v2 := range pairs {
			if !yield(v2, v1) {
				return
			}
		}
	}
}
