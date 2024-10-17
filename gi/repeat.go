package gi

import (
	"iter"

	"golang.org/x/exp/constraints"
)

// RepeatValue returns an iterator that yields the same value repeatedly n times.
func RepeatValue[V any, I constraints.Integer](value V, n I) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := I(0); i < n; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// RepeatPair returns an iterator that yields the same pair of values repeatedly n times.
func RepeatPair[V1, V2 any, I constraints.Integer](v1 V1, v2 V2, n I) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for i := I(0); i < n; i++ {
			if !yield(v1, v2) {
				return
			}
		}
	}
}

// Repeat returns an iterator that yields the same sequence of values repeatedly n times.
func Repeat[V any, I constraints.Integer](seq iter.Seq[V], n I) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := I(0); i < n; i++ {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// RepeatPairs returns an iterator that yields the same sequence of pairs repeatedly n times.
func RepeatPairs[V1, V2 any, I constraints.Integer](seq iter.Seq2[V1, V2], n I) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for i := I(0); i < n; i++ {
			for v1, v2 := range seq {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}
