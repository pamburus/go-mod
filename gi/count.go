package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
)

// Count returns the number of values that match the predicate.
func Count[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) int {
	return Sum(Map(values, func(v V) int {
		if predicate(v) {
			return 1
		}

		return 0
	}))
}

// CountPairs returns the number of pairs that match the predicate.
func CountPairs[V1, V2 any, P gic.PairPredicate[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) int {
	fold := PairFoldFunc(func(v1 V1, v2 V2) int {
		if predicate(v1, v2) {
			return 1
		}

		return 0
	})

	return Sum(fold(pairs))
}
