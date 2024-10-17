package gi

import (
	"iter"
	"slices"
)

// Values returns an iterator over the given values.
func Values[V any](values ...V) iter.Seq[V] {
	switch len(values) {
	case 0:
		return Empty[V]()
	case 1:
		return Single(values[0])
	default:
		return slices.Values(values)
	}
}
