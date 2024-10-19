package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Equal returns a predicate that returns true in case its arguments are equal to the values v1 and v2.
func Equal[V1, V2 comparable](v1 V1, v2 V2) func(V1, V2) bool {
	return func(o1 V1, o2 V2) bool {
		return v1 == o1 && v2 == o2
	}
}

// NotEqual returns a predicate that returns true in case its arguments are not equal to the values v1 and v2.
func NotEqual[V1, V2 comparable](v1 V1, v2 V2) func(V1, V2) bool {
	return Not(Equal(v1, v2))
}

// Less returns a predicate that returns true in case its arguments are less than the values v1 and v2.
// Values are compared in the order they are given.
func Less[V1, V2 constraints.Ordered](o1 V1, o2 V2) func(V1, V2) bool {
	return func(v1 V1, v2 V2) bool {
		switch {
		case v1 != o1:
			return v1 < o1
		default:
			return v2 < o2
		}
	}
}

// LessOrEqual returns a predicate that returns true in case its arguments are less or equal than the values v1 and v2.
// Values are compared in the order they are given.
func LessOrEqual[V1, V2 constraints.Ordered](o1 V1, o2 V2) func(V1, V2) bool {
	return func(v1 V1, v2 V2) bool {
		switch {
		case v1 != o1:
			return v1 <= o1
		default:
			return v2 <= o2
		}
	}
}

// Greater returns a predicate that returns true in case its arguments are greater than the values v1 and v2.
// Values are compared in the order they are given.
func Greater[V1, V2 constraints.Ordered](o1 V1, o2 V2) func(V1, V2) bool {
	return Not(LessOrEqual(o1, o2))
}

// GreaterOrEqual returns a predicate that returns true in case its arguments are greater or equal than the values v1 and v2.
// Values are compared in the order they are given.
func GreaterOrEqual[V1, V2 constraints.Ordered](o1 V1, o2 V2) func(V1, V2) bool {
	return Not(Less(o1, o2))
}

// Not returns a predicate that returns true only in case the given predicate returns false.
func Not[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) P {
	return func(v1 V1, v2 V2) bool {
		return !predicate(v1, v2)
	}
}

// And returns a predicate that returns true only in case all of the given predicates return true.
// And returns true if there are no given predicates.
func And[V1, V2 any, P constraints.Predicate2[V1, V2]](predicates ...P) P {
	return func(v1 V1, v2 V2) bool {
		for _, predicate := range predicates {
			if !predicate(v1, v2) {
				return false
			}
		}

		return true
	}
}

// Or returns a predicate that returns true only in case any of the given predicates return true.
// Or returns false if there are no given predicates.
func Or[V1, V2 any, P constraints.Predicate2[V1, V2]](predicates ...P) P {
	return func(v1 V1, v2 V2) bool {
		for _, predicate := range predicates {
			if predicate(v1, v2) {
				return true
			}
		}

		return false
	}
}

// IsZero returns true if both values are zero.
func IsZero[V1, V2 comparable](v1 V1, v2 V2) bool {
	var z1 V1
	var z2 V2

	return v1 == z1 && v2 == z2
}

// IsNotZero returns true if any of the values is not zero.
func IsNotZero[V1, V2 comparable](v1 V1, v2 V2) bool {
	return !IsZero(v1, v2)
}

// Each returns a predicate that returns true only in case all of the pairs match the predicate.
func Each[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) bool {
	return func(pairs iter.Seq2[V1, V2]) bool {
		return !Contains(pairs, Not(predicate))
	}
}

// Any returns a predicate that returns true in case any of the pairs match the predicate.
func Any[V1, V2 any, P constraints.Predicate2[V1, V2]](predicate P) func(iter.Seq2[V1, V2]) bool {
	return func(pairs iter.Seq2[V1, V2]) bool {
		return Contains(pairs, predicate)
	}
}
