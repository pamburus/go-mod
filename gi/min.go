package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
	"github.com/pamburus/go-mod/gi/giop"
)

// Min returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
func Min[V constraints.Ordered](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.Min)
}

// MinBy returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided key function for comparison.
func MinBy[V any, K constraints.Ordered, F ~func(V) K](values iter.Seq[V], key F) (V, bool) {
	return Reduce(values, giop.MinBy(key))
}

// MinByLess returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the Less method for comparison.
func MinByLess[V constraints.OrderedByLess[V]](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.MinByLess)
}

// MinByLessFunc returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided less function for comparison.
func MinByLessFunc[V any, F ~func(V, V) bool](values iter.Seq[V], less F) (V, bool) {
	return Reduce(values, giop.MinByLessFunc(less))
}

// MinByCompare returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the Compare method for comparison.
func MinByCompare[V constraints.OrderedByCompare[V]](values iter.Seq[V]) (V, bool) {
	return Reduce(values, giop.MinByCompare)
}

// MinByCompareFunc returns the minimum value of the given values and true if values is not empty.
// Otherwise, it returns the zero value of the value type and false.
// It uses the provided compare function for comparison.
func MinByCompareFunc[V any, F ~func(V, V) int](values iter.Seq[V], compare F) (V, bool) {
	return Reduce(values, giop.MinByCompareFunc(compare))
}
