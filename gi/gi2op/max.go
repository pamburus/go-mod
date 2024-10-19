package gi2op

import (
	"github.com/pamburus/go-mod/gi/constraints"
)

// Max returns the maximum value of the left and right.
func Max[V1, V2 constraints.Ordered](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	switch {
	case l1 > r1:
		return l1, l2
	case l1 < r1:
		return r1, r2
	case l2 > r2:
		return l1, l2
	default:
		return r1, r2
	}
}

// MaxBy returns a function that returns the maximum value of the left and right
// using provided key function for comparison.
func MaxBy[V1, V2 any, K constraints.Ordered, Key ~func(V1, V2) K](key Key) func(V1, V2, V1, V2) (V1, V2) {
	return func(l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
		if key(l1, l2) >= key(r1, r2) {
			return l1, l2
		}

		return r1, r2
	}
}

// MaxByLeft returns the pair with the maximum left value.
func MaxByLeft[V1 constraints.Ordered, V2 any](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return MaxBy(func(l1 V1, _ V2) V1 {
		return l1
	})(l1, l2, r1, r2)
}

// MaxByRight returns the pair with the maximum right value.
func MaxByRight[V1 any, V2 constraints.Ordered](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return MaxBy(func(_ V1, l2 V2) V2 {
		return l2
	})(l1, l2, r1, r2)
}
