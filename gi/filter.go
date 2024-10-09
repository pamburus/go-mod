package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
)

// Filter returns an iterator adapter that uses a predicate to filter values.
func Filter[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) iter.Seq[V] {
	return func(yield func(V) bool) {
		for value := range values {
			if predicate(value) {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// FilterPairs returns an iterator adapter over pairs that uses a predicate to filter pairs.
func FilterPairs[V1, V2 any, P ~func(V1, V2) bool](pairs iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
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

// FilterLeft returns an iterator adapter over pairs that uses a predicate on the first value of a pair to filter pairs.
func FilterLeft[V1, V2 any, P gic.Predicate[V1]](pairs iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
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

// FilterRight returns an iterator adapter over pairs that uses a predicate on the second value of a pair to filter pairs.
func FilterRight[V1, V2 any, P gic.Predicate[V2]](values iter.Seq2[V1, V2], predicate P) iter.Seq2[V1, V2] {
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
func FilterKeys[K, V any, P gic.Predicate[K]](pairs iter.Seq2[K, V], predicate P) iter.Seq2[K, V] {
	return FilterLeft(pairs, predicate)
}

// FilterValues is an alias for [FilterRight] that assumes the given iterator is an iterator over key/value pairs.
func FilterValues[K, V any, P gic.Predicate[V]](pairs iter.Seq2[K, V], predicate P) iter.Seq2[K, V] {
	return FilterRight(pairs, predicate)
}
