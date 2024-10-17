package gi2

import "iter"

// Loop returns an iterator that yields the same sequence of pairs repeatedly.
func Loop[V1, V2 any](seq iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for {
			for v1, v2 := range seq {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}

// LoopSingle returns an iterator that yields the same pair of values repeatedly.
func LoopSingle[V1, V2 any](v1 V1, v2 V2) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for {
			if !yield(v1, v2) {
				return
			}
		}
	}
}
