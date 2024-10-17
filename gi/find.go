package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/optional/optval"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) optval.Value[V] {
	for v := range values {
		if predicate(v) {
			return optval.Some(v)
		}
	}

	return optval.None[V]()
}
