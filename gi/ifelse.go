package gi

import "github.com/pamburus/go-mod/gi/constraints"

func IfElse[R any](condition bool, tv, fv R) R {
	if condition {
		return tv
	}

	return fv
}

func IfElseFunc[T any, R any, P constraints.Predicate[T]](predicate P, tv, fv R) func(T) R {
	return func(value T) R {
		return IfElse(predicate(value), tv, fv)
	}
}
