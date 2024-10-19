package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Filter returns an iterator adapter over pairs that uses a predicate to filter them.
func Filter[V1, V2 any, P constraints.Predicate2[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
	return FilterWith(predicate)(pairs)
}

// FilterWith returns a function that filters an iterator sequence over pairs using a predicate.
func FilterWith[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return func(pairs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
		return func(yield func(V1, V2) bool) {
			for v1, v2 := range pairs {
				if predicate(v1, v2) {
					if !yield(v1, v2) {
						return
					}
				}
			}
		}
	}
}

// FilterLeft returns an iterator adapter over pairs that uses a predicate on the first value of a pair to filter them.
func FilterLeft[V1, V2 any, P constraints.Predicate[V1]](pairs iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for v1, v2 := range pairs {
			if predicate(v1) {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}

// FilterRight returns an iterator adapter over pairs that uses a predicate on the second value of a pair to filter them.
func FilterRight[V1, V2 any, P constraints.Predicate[V2]](values iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		for v1, v2 := range values {
			if predicate(v2) {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}

// FilterKeys is an alias for [FilterLeft] that assumes the given iterator is an iterator over key/value pairs.
func FilterKeys[K, V any, P constraints.Predicate[K]](pairs iter.Seq2[K, V], predicate P) iter.Seq2[K, V] {
	return FilterLeft(pairs, predicate)
}

// FilterValues is an alias for [FilterRight] that assumes the given iterator is an iterator over key/value pairs.
func FilterValues[K, V any, P constraints.Predicate[V]](pairs iter.Seq2[K, V], predicate P) iter.Seq2[K, V] {
	return FilterRight(pairs, predicate)
}
