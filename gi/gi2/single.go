package gi2

import "iter"

// Single returns an iterator over a single pair of values.
func Single[V1, V2 any](v1 V1, v2 V2) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		yield(v1, v2)
	}
}
