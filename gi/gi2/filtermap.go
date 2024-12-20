package gi2

import (
	"iter"
)

// FilterMap returns an iterator that yields the results of applying f to each pair of values.
// If f returns false for the pair, it is skipped.
func FilterMap[V1, V2, R1, R2 any, F func(V1, V2) (R1, R2, bool)](pairs iter.Seq2[V1, V2], f F) iter.Seq2[R1, R2] {
	return FilterMapWith(f)(pairs)
}

// FilterMapWith returns a function that transforms an iterator sequence over pairs to an iterator sequence over pairs.
// It uses the provided filter-map function to transform and filter pairs.
// If f returns false for the pair, it is skipped.
func FilterMapWith[V1, V2, R1, R2 any, F func(V1, V2) (R1, R2, bool)](f F) func(pairs iter.Seq2[V1, V2]) iter.Seq2[R1, R2] {
	return func(pairs iter.Seq2[V1, V2]) iter.Seq2[R1, R2] {
		return func(yield func(R1, R2) bool) {
			for v1, v2 := range pairs {
				if r1, r2, ok := f(v1, v2); ok {
					if !yield(r1, r2) {
						return
					}
				}
			}
		}
	}
}
