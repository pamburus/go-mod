package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Every returns true if all pairs match the predicate.
func Every[V1, V2 any, P constraints.Predicate2[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) bool {
	return !Contains(pairs, Not(predicate))
}
