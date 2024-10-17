package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
	"github.com/pamburus/go-mod/optional/optpair"
	"github.com/pamburus/go-mod/optional/optval"
)

// Find returns the first value among the values matching the predicate.
func Find[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) (V, bool) {
	for v := range values {
		if predicate(v) {
			return v, true
		}
	}

	var zero V

	return zero, false
}

// FindOpt returns the first value among the values matching the predicate.
func FindOpt[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) optval.Value[V] {
	return optval.New(Find(values, predicate))
}

// FindWith returns a function that returns the first value among the values matching the predicate.
func FindWith[V any, P gic.Predicate[V]](predicate P) func(iter.Seq[V]) (V, bool) {
	return func(values iter.Seq[V]) (V, bool) {
		return Find(values, predicate)
	}
}

// FindOptWith returns a function that returns the first value among the values matching the predicate.
func FindOptWith[V any, P gic.Predicate[V]](predicate P) func(iter.Seq[V]) optval.Value[V] {
	return func(values iter.Seq[V]) optval.Value[V] {
		return FindOpt(values, predicate)
	}
}

// FindPair returns the first pair among the pairs matching the predicate.
func FindPair[V1, V2 any, P gic.PairPredicate[V1, V2]](values iter.Seq2[V1, V2], predicate P) (V1, V2, bool) {
	for v1, v2 := range values {
		if predicate(v1, v2) {
			return v1, v2, true
		}
	}

	var z1 V1
	var z2 V2

	return z1, z2, false
}

// FindPairOpt returns the first pair among the pairs matching the predicate.
func FindPairOpt[V1, V2 any, P gic.PairPredicate[V1, V2]](values iter.Seq2[V1, V2], predicate P) optpair.Pair[V1, V2] {
	return optpair.New(FindPair(values, predicate))
}

// FindPairWith returns a function that returns the first pair among the pairs matching the predicate.
func FindPairWith[V1, V2 any, P gic.PairPredicate[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) (V1, V2, bool) {
	return func(values iter.Seq2[V1, V2]) (V1, V2, bool) {
		return FindPair(values, predicate)
	}
}

// FindPairOptWith returns a function that returns the first pair among the pairs matching the predicate.
func FindPairOptWith[V1, V2 any, P gic.PairPredicate[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
	return func(values iter.Seq2[V1, V2]) optpair.Pair[V1, V2] {
		return FindPairOpt(values, predicate)
	}
}
