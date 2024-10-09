// Package gic provides constraints for generic iterator helpers.
package gic

import "golang.org/x/exp/constraints"

// Cloneable is a type that can be cloned.
type Cloneable[T any] interface {
	Clone() T
}

// Predicate is a function that takes an argument of type T and returns a bool.
type Predicate[T any] interface {
	~func(T) bool
}

// PairPredicate is a function that takes two arguments of types T1 and T2 and returns a bool.
type PairPredicate[T1, T2 any] interface {
	~func(T1, T2) bool
}

// Mapping is a type based on a map with keys of type K and values of type V.
type Mapping[K comparable, V any] interface {
	~map[K]V
}

// OrderedByLess is a type that can be ordered by the Less method.
type OrderedByLess[T any] interface {
	Less(T) bool
}

// OrderedByCompare is a type that can be ordered by the Compare method.
type OrderedByCompare[T any] interface {
	Compare(T) int
}

// ComparableByEqual is a type that can be compared by the Equal method.
type ComparableByEqual[T any] interface {
	Equal(T) bool
}

// Number is a type that can be either an integer or a float.
type Number interface {
	constraints.Integer | constraints.Float
}
