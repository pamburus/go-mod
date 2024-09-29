package gi

import "iter"

// Limit returns an iterator sequence that yields at most n values.
func Limit[V any](values iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		for v := range values {
			if i >= n {
				return
			}

			if !yield(v) {
				return
			}

			i++
		}
	}
}

// LimitPairs returns an iterator sequence that yields at most n pairs.
func LimitPairs[V1, V2 any](pairs iter.Seq2[V1, V2], n int) iter.Seq2[V1, V2] {
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
