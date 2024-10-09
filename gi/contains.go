package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
)

// Contains returns true if the given values contain a value matching the predicate.
func Contains[V any, P gic.Predicate[V]](values iter.Seq[V], predicate P) bool {
	return Find(values, predicate).IsSome()
}

// ContainsPair returns true if the given pairs contain a pair matching the predicate.
func ContainsPair[V1, V2 any, P gic.PairPredicate[V1, V2]](pairs iter.Seq2[V1, V2], predicate P) bool {
	return FindPair(pairs, predicate).IsSome()
}

// ContainsLeft returns true if the given pairs contain a pair with the first value matching the predicate.
func ContainsLeft[V1 comparable, V2 any, P gic.Predicate[V1]](pairs iter.Seq2[V1, V2], predicate P) bool {
	return ContainsPair(pairs, func(v1 V1, _ V2) bool {
		return predicate(v1)
	})
}

// ContainsRight returns true if the given pairs contain a pair with the second value matching the predicate.
func ContainsRight[V1 any, V2 comparable, P gic.Predicate[V2]](pairs iter.Seq2[V1, V2], predicate P) bool {
	return ContainsPair(pairs, func(_ V1, v2 V2) bool {
		return predicate(v2)
	})
}

// ContainsKey is an alias for [ContainsLeft] that assumes the given iterator is an iterator over key/value pairs.
func ContainsKey[K comparable, V any, P gic.Predicate[K]](pairs iter.Seq2[K, V], predicate P) bool {
	return ContainsLeft(pairs, predicate)
}

// ContainsValue is an alias for [ContainsRight] that assumes the given iterator is an iterator over key/value pairs.
func ContainsValue[K any, V comparable, P gic.Predicate[V]](pairs iter.Seq2[K, V], predicate P) bool {
	return ContainsRight(pairs, predicate)
}
