package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
)

// Count returns the number of pairs that match the predicate.
func Count[V1, V2 any, P constraints.Predicate2[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) int {
	return gi.Sum(PairFold(pairs, IfElseFunc(predicate, 1, 0)))
}
