package gi

import (
	"iter"
)

// Single returns an iterator over a single value.
func Single[V any](value V) iter.Seq[V] {
	return func(yield func(V) bool) {
		yield(value)
	}
}
