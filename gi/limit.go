package gi

import "iter"

// Limit returns an iterator sequence that yields at most n values.
func Limit[V any](values iter.Seq[V], n int) iter.Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		for v := range values {
			if i >= n {
				return
			}

			if !yield(v) {
				return
			}

			i++
		}
	}
}
