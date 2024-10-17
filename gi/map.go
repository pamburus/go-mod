package gi

import "iter"

// Map returns an iterator adapter that uses a transform function to transform values.
func Map[V, R any, F ~func(V) R](values iter.Seq[V], transform F) iter.Seq[R] {
	return func(yield func(R) bool) {
		for value := range values {
			if !yield(transform(value)) {
				return
			}
		}
	}
}
