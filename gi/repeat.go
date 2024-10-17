package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

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

// RepeatSingle returns an iterator that yields the same single value repeatedly n times.
func RepeatSingle[V any, I constraints.Integer](value V, n I) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := I(0); i < n; i++ {
			if !yield(value) {
				return
			}
		}
	}
}
