package gi

import "iter"

// FilterMap returns an iterator that yields values by applying the provided filter-map function to values.
// The filter-map function returns a value and a boolean indicating whether the value should be included in the output.
func FilterMap[V, R any, F func(V) (R, bool)](values iter.Seq[V], f F) iter.Seq[R] {
	return FilterMapWith(f)(values)
}

// FilterMapWith returns a function that transforms an iterator sequence over values to an iterator sequence over values.
// It uses the provided filter-map function to transform values to values.
// If f returns false, the value is skipped.
func FilterMapWith[V, R any, F func(V) (R, bool)](f F) func(values iter.Seq[V]) iter.Seq[R] {
	return func(values iter.Seq[V]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for v := range values {
				if r, ok := f(v); ok {
					if !yield(r) {
						return
					}
				}
			}
		}
	}
}

// FilterMapPairs returns an iterator sequence that yields the results of applying f to each pair of values.
// If f returns false, the pair is skipped.
func FilterMapPairs[V1, V2, R1, R2 any, F func(V1, V2) (R1, R2, bool)](pairs iter.Seq2[V1, V2], f F) iter.Seq2[R1, R2] {
	return FilterMapPairsWith(f)(pairs)
}

// FilterMapPairsWith returns a function that transforms an iterator sequence over pairs to an iterator sequence over pairs.
// It uses the provided filter-map function to transform pairs to pairs.
// If f returns false, the pair is skipped.
func FilterMapPairsWith[V1, V2, R1, R2 any, F func(V1, V2) (R1, R2, bool)](f F) func(pairs iter.Seq2[V1, V2]) iter.Seq2[R1, R2] {
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
