package gi2

import "iter"

// Limit returns an iterator sequence that yields at most n pairs.
func Limit[V1, V2 any](pairs iter.Seq2[V1, V2], n int) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		i := 0
		for v1, v2 := range pairs {
			if i >= n {
				return
			}

			if !yield(v1, v2) {
				return
			}

			i++
		}
	}
}
