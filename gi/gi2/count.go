package gi2

import (
	"iter"
)

// Count returns the number of pairs that match the predicate.
func Count[V1, V2 any](pairs iter.Seq2[V1, V2]) int {
	n := 0

	for range pairs {
		n++
	}

	return n
}
