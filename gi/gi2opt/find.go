package gi2opt

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/optional/optpair"
)

// Find returns the first pair among the pairs matching the predicate.
func Find[V1, V2 any, P constraints.Predicate2[V1, V2]](values iter.Seq2[V1, V2], predicate P) optpair.Pair[V1, V2] {
	return optpair.New(gi2.Find(values, predicate))
}
