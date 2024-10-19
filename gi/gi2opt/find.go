package gi2opt

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Find returns the first pair among the pairs matching the predicate.
func Find[V1, V2 any, P constraints.Predicate2[V1, V2]](values iter.Seq2[V1, V2], predicate P) optpair.Pair[V1, V2] {
	return FindWith(predicate)(values)
}

// FindWith returns a function that finds the first pair among the pairs matching the predicate.
func FindWith[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return func(values iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
		return optpair.New(gi2.Find(values, predicate))
	}
}
