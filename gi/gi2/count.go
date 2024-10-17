package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
)

// Count returns the number of pairs that match the predicate.
func Count[V1, V2 any, P constraints.Predicate2[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) int {
	fold := PairFoldWith(func(v1 V1, v2 V2) int {
		if predicate(v1, v2) {
			return 1
		}

		return 0
	})

	return gi.Sum(fold(pairs))
}
