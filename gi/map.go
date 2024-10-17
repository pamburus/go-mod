package gi

import "iter"

// Map returns an iterator adapter that uses a transform function to transform values.
func Map[V, R any, F ~func(V) R](values iter.Seq[V], transform F) iter.Seq[R] {
	return MapWith(transform)(values)
}

// MapWith returns a function that transforms an iterator sequence over values to an iterator sequence over transformed values.
func MapWith[V, R any, F ~func(V) R](transform F) func(iter.Seq[V]) iter.Seq[R] {
	return func(values iter.Seq[V]) iter.Seq[R] {
		return func(yield func(R) bool) {
			for value := range values {
				if !yield(transform(value)) {
					return
				}
			}
		}
	}
}
