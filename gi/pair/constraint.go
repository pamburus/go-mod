package pair

// Predicate is a function taking arguments of type T1 and T2 and returning a bool based on the values.
type Predicate[T1, T2 any] interface {
	~func(T1, T2) bool
}
