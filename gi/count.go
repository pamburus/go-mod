package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Count returns the number of values that match the predicate.
func Count[V any, P constraints.Predicate[V]](values iter.Seq[V], predicate P) int {
	return Sum(Map(values, IfElseFunc(predicate, 1, 0)))
}
