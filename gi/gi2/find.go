package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Find returns the first pair among the pairs matching the predicate.
func Find[V1, V2 any, P constraints.Predicate2[V1, V2]](values iter.Seq2[V1, V2], predicate P) (V1, V2, bool) {
	return FindWith(predicate)(values)
}

// FindWith returns a function that finds the first pair among the pairs matching the predicate.
func FindWith[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) (V1, V2, bool) {
	return func(values iter.Seq2[V1, V2]) (V1, V2, bool) {
		for v1, v2 := range values {
			if predicate(v1, v2) {
				return v1, v2, true
			}
		}

		return zero[V1](), zero[V2](), false
	}
}
