package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
	"github.com/pamburus/go-mod/gi/pair"
)

// Every returns true if all values match the predicate.
func Every[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) bool {
	return !Contains(values, Not(predicate))
}

// EveryPair returns true if all pairs match the predicate.
func EveryPair[V1, V2 any, P gic.PairPredicate[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) bool {
	return !ContainsPair(pairs, pair.Not(predicate))
}
