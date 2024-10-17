package gi

import "iter"

// FilterPairFold returns an iterator sequence that yields the results of applying f to each pair of values.
// If fold returns false, the pair is skipped.
func FilterPairFold[V1, V2, R any, F func(V1, V2) (R, bool)](pairs iter.Seq2[V1, V2], fold F) iter.Seq[R] {
	return FilterPairFoldWith(fold)(pairs)
}

// FilterPairFoldWith returns a function that transforms an iterator sequence over pairs to an iterator sequence over values.
// It uses the provided fold function to transform pairs to values.
// If fold returns false, the pair is skipped.
func FilterPairFoldWith[V1, V2, R any, F func(V1, V2) (R, bool)](fold F) func(iter.Seq2[V1, V2]) iter.Seq[R] {
	return func(pairs iter.Seq2[V1, V2]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v1, v2 := range pairs {
				if r, ok := fold(v1, v2); ok {
					if !yield(r) {
						return
					}
				}
			}
		}
	}
}

// FilterPairUnfold returns an iterator sequence that yields the results of applying f to each value.
// If unfold returns false, the value is skipped.
func FilterPairUnfold[V, R1, R2 any, F func(V) (R1, R2, bool)](values iter.Seq[V], unfold F) iter.Seq2[R1, R2] {
	return FilterPairUnfoldWith(unfold)(values)
}

// FilterPairUnfoldWith returns a function that transforms an iterator sequence over values to an iterator sequence over pairs.
// It uses the provided unfold function to transform values to pairs.
// If unfold returns false, the value is skipped.
func FilterPairUnfoldWith[V, R1, R2 any, F func(V) (R1, R2, bool)](unfold F) func(iter.Seq[V]) iter.Seq2[R1, R2] {
	return func(values iter.Seq[V]) iter.Seq2[R1, R2] {
		return func(yield func(R1, R2) bool) {
			for v := range values {
				if r1, r2, ok := unfold(v); ok {
					if !yield(r1, r2) {
						return
					}
				}
			}
		}
	}
}
