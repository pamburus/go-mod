package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/giop"
)

// Max returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
func Max[V constraints.Ordered](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.Max)
}

// MaxBy returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided key function for comparison.
func MaxBy[V any, K constraints.Ordered, F ~func(V) K](values iter.Seq[V], key F) (V, bool) {
	return Reduce(values, giop.MaxBy(key))
}

// MaxByLess returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the Less method for comparison.
func MaxByLess[V constraints.OrderedByLess[V]](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.MaxByLess)
}

// MaxByLessFunc returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided less function for comparison.
func MaxByLessFunc[V any, F ~func(V, V) bool](values iter.Seq[V], less F) (V, bool) {
	return Reduce(values, giop.MaxByLessFunc(less))
}

// MaxByCompare returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the Compare method for comparison.
func MaxByCompare[V constraints.OrderedByCompare[V]](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.MaxByCompare)
}

// MaxByCompareFunc returns the maximum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided compare function for comparison.
func MaxByCompareFunc[V any, F ~func(V, V) int](values iter.Seq[V], compare F) (V, bool) {
	return Reduce(values, giop.MaxByCompareFunc(compare))
}
