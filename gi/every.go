package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Every returns true if all values match the predicate.
func Every[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) bool {
	return !Contains(values, Not(predicate))
}
