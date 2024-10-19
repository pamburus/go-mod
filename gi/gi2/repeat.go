package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Repeat returns an iterator that yields the same sequence of pairs repeatedly n times.
func Repeat[V1, V2 any, I constraints.Integer](seq iter.Seq2[V1, V2], n I) iter.Seq2[V1, V2] {
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

// RepeatSingle returns an iterator that yields the same single pair of values repeatedly n times.
func RepeatSingle[V1, V2 any, I constraints.Integer](v1 V1, v2 V2, n I) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for i := I(0); i < n; i++ {
			if !yield(v1, v2) {
				return
			}
		}
	}
}
