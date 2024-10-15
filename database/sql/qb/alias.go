package qb

type Aliased[T any] struct {
	Alias string
	Value T
}
