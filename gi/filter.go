package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Filter returns an iterator adapter that uses a predicate to filter values.
func Filter[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) iter.Seq[V] {
	return FilterWith(predicate)(values)
}

// FilterWith returns a function that filters an iterator sequence over values using a predicate.
func FilterWith[V any, P constraints.Predicate[V]](predicate P) func(iter.Seq[V]) iter.Seq[V] {
	return func(values iter.Seq[V]) iter.Seq[V] {
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
}
