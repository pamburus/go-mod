package gi2

import "iter"

// Left returns an iterator that yields left value from each pair and discards right value.
func Left[V1, V2 any](pairs iter.Seq2[V1, V2]) iter.Seq[V1] {
	return func(yield func(V1) bool) {
		for v1 := range pairs {
			if !yield(v1) {
				return
			}
		}
	}
}

// Right returns an iterator that yields right value from each pair and discards left value.
func Right[V1, V2 any](pairs iter.Seq2[V1, V2]) iter.Seq[V2] {
	return func(yield func(V2) bool) {
		for _, v2 := range pairs {
			if !yield(v2) {
				return
			}
		}
	}
}
