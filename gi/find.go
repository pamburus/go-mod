package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) (V, bool) {
	for v := range values {
		if predicate(v) {
			return v, true
		}
	}

	return zero[V](), false
}
