package gi

import (
	"iter"
)

// LoopValue returns an iterator that yields the same value repeatedly.
func LoopValue[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			if !yield(value) {
				return
			}
		}
	}
}

// LoopPair returns an iterator that yields the same pair of values repeatedly.
func LoopPair[V1, V2 any](v1 V1, v2 V2) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for {
			if !yield(v1, v2) {
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

// LoopPairs returns an iterator that yields the same sequence of pairs repeatedly.
func LoopPairs[V1, V2 any](seq iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
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
