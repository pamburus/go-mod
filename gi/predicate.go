package gi

import (
	"slices"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Equal returns a predicate that returns true in case its argument is equal to the value.
func Equal[T comparable](value T) func(T) bool {
	return func(other T) bool {
		return other == value
	}
}

// NotEqual returns a predicate that returns true in case its argument is not equal to the value.
func NotEqual[T comparable](value T) func(T) bool {
	return Not(Equal(value))
}

// Less returns a predicate that returns true in case its argument is less than the value.
func Less[T constraints.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return other < value
	}
}

// LessOrEqual returns a predicate that returns true in case its argument is less or equal than the value.
func LessOrEqual[T constraints.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return other <= value
	}
}

// Greater returns a predicate that returns true in case its argument is greater than the value.
func Greater[T constraints.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return other > value
	}
}

// GreaterOrEqual returns a predicate that returns true in case its argument is greater or equal than the value.
func GreaterOrEqual[T constraints.Ordered](value T) func(T) bool {
	return func(other T) bool {
		return other >= value
	}
}

// Even returns true if the value can be divided by 2 without a remainder.
func Even[T constraints.Integer](value T) bool {
	return value%2 == 0
}

// Odd returns true if the value can not be divided by 2 without a remainder.
func Odd[T constraints.Integer](value T) bool {
	return value%2 == 1
}

// DivisibleBy returns a predicate that returns true if its argument can be divided by the given divisor without a remainder.
func DivisibleBy[T constraints.Integer](divisor T) func(value T) bool {
	return func(value T) bool {
		return value%divisor == 0
	}
}

// Not returns a predicate that returns negated value of the given predicate.
func Not[T any, P constraints.Predicate[T]](predicate P) P {
	return func(value T) bool {
		return !predicate(value)
	}
}

// And returns a predicate that returns true only in case all of the given predicates return true.
// And returns true if there are no given predicates.
func And[T any, P constraints.Predicate[T]](predicates ...P) P {
	return func(value T) bool {
		for _, predicate := range predicates {
			if !predicate(value) {
				return false
			}
		}

		return true
	}
}

// Or returns a predicate that returns true only in case any of the given predicates return true.
// Or returns false if there are no given predicates.
func Or[T any, P constraints.Predicate[T]](predicates ...P) P {
	return func(value T) bool {
		for _, predicate := range predicates {
			if predicate(value) {
				return true
			}
		}

		return false
	}
}

// In returns a predicate that returns true in case its argument is contained in the given slice of values.
func In[T comparable, TT ~[]T](values TT) func(T) bool {
	return func(value T) bool {
		return Contains(slices.Values(values), Equal(value))
	}
}

// OneOf returns a predicate that returns true in case its argument is equal to one of the given values.
func OneOf[T comparable](values ...T) func(T) bool {
	return In(values)
}

// IsZero returns true if the value is equal to the zero value of its type.
func IsZero[T comparable](value T) bool {
	var zero T

	return value == zero
}

// IsNotZero returns true if the value is not equal to the zero value of its type.
func IsNotZero[T comparable](value T) bool {
	return !IsZero(value)
}
