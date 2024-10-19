package gi

import "iter"

// LoopSingle returns an iterator that yields the same value repeatedly.
func LoopSingle[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}

// Loop returns an iterator that yields the same sequence of values repeatedly.
func Loop[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}
