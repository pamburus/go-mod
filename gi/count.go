package gi

import (
	"iter"
)

// Count returns the number of values in the sequence.
func Count[V any](values iter.Seq[V]) int {
	n := 0

	for range values {
		n++
	}

	return n
}
