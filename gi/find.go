package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
	"github.com/pamburus/go-mod/optional/optpair"
	"github.com/pamburus/go-mod/optional/optval"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) optval.Value[V] {
	for v := range values {
		if predicate(v) {
			return optval.Some(v)
		}
	}

	return optval.None[V]()
}

// FindPair returns the first pair among the pairs matching the predicate.
func FindPair[V1, V2 any, P gic.PairPredicate[V1, V2]](values iter.Seq2[V1, V2], predicate P) optpair.Pair[V1, V2] {
	for v1, v2 := range values {
		if predicate(v1, v2) {
			return optpair.Some(v1, v2)
		}
	}

	return optpair.None[V1, V2]()
}
