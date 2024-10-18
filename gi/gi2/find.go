package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Find returns the first pair among the pairs matching the predicate.
func Find[V1, V2 any, P constraints.Predicate2[V1, V2]](values iter.Seq2[V1, V2], predicate P) optpair.Pair[V1, V2] {
	for v1, v2 := range values {
		if predicate(v1, v2) {
			return optpair.Some(v1, v2)
		}
	}

	return optpair.None[V1, V2]()
}
