package gi2

import "iter"

// Swap returns an iterator that yields pairs with swapped values.
func Swap[V1, V2 any](pairs iter.Seq2[V1, V2]) iter.Seq2[V2, V1] {
	return func(yield func(V2, V1) bool) {
		for v1, v2 := range pairs {
			if !yield(v2, v1) {
				return
			}
		}
	}
}
