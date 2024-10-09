package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/gic"
)

// Cloned returns a new iterator that consumes original iterator and clones each value.
func Cloned[V gic.Cloneable[V]](values iter.Seq[V]) iter.Seq[V] {
	return Map(values, V.Clone)
}

// ClonedPairs returns a new iterator that consumes original iterator and clones each value in the pair.
func ClonedPairs[V1 gic.Cloneable[V1], V2 gic.Cloneable[V2]](pairs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return MapPairs(pairs, func(v1 V1, v2 V2) (V1, V2) {
		return v1.Clone(), v2.Clone()
	})
}

// ClonedLeft returns a new iterator that consumes original iterator and clones each first value in the pair.
func ClonedLeft[V1 gic.Cloneable[V1], V2 any](pairs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return MapLeft(pairs, V1.Clone)
}

// ClonedRight returns a new iterator that consumes original iterator and clones each second value in the pair.
func ClonedRight[V1 any, V2 gic.Cloneable[V2]](pairs iter.Seq2[V1, V2]) iter.Seq2[V1, V2] {
	return MapRight(pairs, V2.Clone)
}

// ClonedKeys is an alias for [ClonedLeft] that assumes the given iterator is an iterator over key/value pairs.
func ClonedKeys[K gic.Cloneable[K], V any](pairs iter.Seq2[K, V]) iter.Seq2[K, V] {
	return ClonedLeft(pairs)
}

// ClonedValues is an alias for [ClonedRight] that assumes the given iterator is an iterator over key/value pairs.
func ClonedValues[K any, V gic.Cloneable[V]](pairs iter.Seq2[K, V]) iter.Seq2[K, V] {
	return ClonedRight(pairs)
}
