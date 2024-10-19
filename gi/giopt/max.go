package giopt

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/giop"
	"github.com/pamburus/go-mod/optional/optval"
)

// Max returns the maximum value of the given values or [optval.None] if values is empty.
func Max[V constraints.Ordered](values iter.Seq[V]) optval.Value[V] {
	return Reduce(values, giop.Max)
}

// MaxBy returns the maximum value of the given values or [optval.None] if values is empty.
// It uses provided key function for comparison.
func MaxBy[V any, K constraints.Ordered, F ~func(V) K](values iter.Seq[V], key F) optval.Value[V] {
	return Reduce(values, giop.MaxBy(key))
}

// MaxByLess returns the maximum value of the given values or [optval.None] if values is empty.
// It uses Less method for comparison.
func MaxByLess[V constraints.OrderedByLess[V]](values iter.Seq[V]) optval.Value[V] {
	return Reduce(values, giop.MaxByLess)
}

// MaxByLessFunc returns the maximum value of the given values or [optval.None] if values is empty.
// It uses provided less function for comparison.
func MaxByLessFunc[V any, F ~func(V, V) bool](values iter.Seq[V], less F) optval.Value[V] {
	return Reduce(values, giop.MaxByLessFunc(less))
}

// MaxByCompare returns the maximum value of the given values or [optval.None] if values is empty.
// It uses Compare method for comparison.
func MaxByCompare[V constraints.OrderedByCompare[V]](values iter.Seq[V]) optval.Value[V] {
	return Reduce(values, giop.MaxByCompare)
}

// MaxByCompareFunc returns the maximum value of the given values or [optval.None] if values is empty.
// It uses provided compare function for comparison.
func MaxByCompareFunc[V any, F ~func(V, V) int](values iter.Seq[V], compare F) optval.Value[V] {
	return Reduce(values, giop.MaxByCompareFunc(compare))
}
