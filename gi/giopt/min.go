package giopt

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/optional/optval"
)

// Min returns the minimum value of the given values or [optval.None] if values is empty.
func Min[V constraints.Ordered](values iter.Seq[V]) optval.Value[V] {
	return optval.New(gi.Min(values))
}

// MinBy returns the minimum value of the given values or [optval.None] value if values is empty.
// It uses provided key function for comparison.
func MinBy[V any, K constraints.Ordered, F ~func(V) K](values iter.Seq[V], key F) optval.Value[V] {
	return optval.New(gi.MinBy(values, key))
}

// MinByLess returns the minimum value of the given values or [optval.None] value if values is empty.
// It uses Less method for comparison.
func MinByLess[V constraints.OrderedByLess[V]](values iter.Seq[V]) optval.Value[V] {
	return optval.New(gi.MinByLess(values))
}

// MinByLessFunc returns the minimum value of the given values or [optval.None] value if values is empty.
// It uses provided less function for comparison.
func MinByLessFunc[V any, F ~func(V, V) bool](values iter.Seq[V], less F) optval.Value[V] {
	return optval.New(gi.MinByLessFunc(values, less))
}

// MinByCompare returns the minimum value of the given values or [optval.None] value if values is empty.
// It uses Compare method for comparison.
func MinByCompare[V constraints.OrderedByCompare[V]](values iter.Seq[V]) optval.Value[V] {
	return optval.New(gi.MinByCompare(values))
}

// MinByCompareFunc returns the minimum value of the given values or [optval.None] value if values is empty.
// It uses provided compare function for comparison.
func MinByCompareFunc[V any, F ~func(V, V) int](values iter.Seq[V], compare F) optval.Value[V] {
	return optval.New(gi.MinByCompareFunc(values, compare))
}
