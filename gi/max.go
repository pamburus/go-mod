package gi

import (
	"iter"

	"golang.org/x/exp/constraints"

	"github.com/pamburus/go-mod/gi/gic"
	"github.com/pamburus/go-mod/gi/giop"
	"github.com/pamburus/go-mod/optional"
)

// Max returns maximum value of the given values or [optional.None] if values is empty.
func Max[V constraints.Ordered](values iter.Seq[V]) optional.Value[V] {
	return Reduce(values, giop.Max[V])
}

// MaxBy returns maximum value of the given values or [optional.None] if values is empty.
// It uses provided key function for comparison.
func MaxBy[V any, K constraints.Ordered, F ~func(V) K](values iter.Seq[V], key F) optional.Value[V] {
	return Reduce(values, giop.MaxBy(key))
}

// MaxByLess returns maximum value of the given values or [optional.None] if values is empty.
// It uses Less method for comparison.
func MaxByLess[V gic.OrderedByLess[V]](values iter.Seq[V]) optional.Value[V] {
	return Reduce(values, giop.MaxByLess)
}

// MaxByLessFunc returns maximum value of the given values or [optional.None] if values is empty.
// It uses provided less function for comparison.
func MaxByLessFunc[V any, F ~func(V, V) bool](values iter.Seq[V], less F) optional.Value[V] {
	return Reduce(values, giop.MaxByLessFunc(less))
}

// MaxByCompare returns maximum value of the given values or [optional.None] if values is empty.
// It uses Compare method for comparison.
func MaxByCompare[V gic.OrderedByCompare[V]](values iter.Seq[V]) optional.Value[V] {
	return Reduce(values, giop.MaxByCompare)
}

// MaxByCompareFunc returns maximum value of the given values or [optional.None] if values is empty.
// It uses provided compare function for comparison.
func MaxByCompareFunc[V any, F ~func(V, V) int](values iter.Seq[V], compare F) optional.Value[V] {
	return Reduce(values, giop.MaxByCompareFunc(compare))
}
