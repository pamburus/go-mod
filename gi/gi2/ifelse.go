package gi2

import (
	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
)

func IfElseFunc[V1, V2, R any, P constraints.Predicate2[V1, V2]](predicate P, tv, fv R) func(V1, V2) R {
	return func(v1 V1, v2 V2) R {
		return gi.IfElse(predicate(v1, v2), tv, fv)
	}
}
