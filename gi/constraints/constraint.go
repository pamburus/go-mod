// Package constraints provides constraints for generic iterator helpers.
package constraints

// Signed is any signed integer type.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is any unsigned integer type.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is any integer type.
type Integer interface {
	Signed | Unsigned
}

// Float is any floating-point type.
type Float interface {
	~float32 | ~float64
}

// Number is any number type.
type Number interface {
	Integer | Float
}

// Ordered is any built-in type that can be ordered.
type Ordered interface {
	Number | ~string
}

// Predicate is a function that takes an argument of type T and returns a bool.
type Predicate[T any] interface {
	~func(T) bool
}

// Predicate2 is a function that takes arguments of type T1 and T2 and returns a bool.
type Predicate2[T1, T2 any] interface {
	~func(T1, T2) bool
}

// Cloneable is a type that can be cloned.
type Cloneable[T any] interface {
	Clone() T
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
