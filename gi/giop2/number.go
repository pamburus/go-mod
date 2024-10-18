package giop2

import (
	"math"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Add returns sum of l1 and r1 and sum of l2 and r2.
func Add[V1, V2 constraints.Number](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 + r1, l2 + r2
}

// Subtract returns difference of l1 and r1 and difference of l2 and r2.
func Subtract[V1, V2 constraints.Number](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 - r1, l2 - r2
}

// Multiply returns product of l1 and r1 and product of l2 and r2.
func Multiply[V1, V2 constraints.Number](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 * r1, l2 * r2
}

// Divide returns quotient of l1 and r1 and quotient of l2 and r2.
func Divide[V1, V2 constraints.Number](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 / r1, l2 / r2
}

// IntMod returns integer remainder of l1 and r1 and integer remainder of l2 and r2.
func IntMod[V1, V2 constraints.Integer](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 % r1, l2 % r2
}

// Mod returns math.Mod(l1, r1) and math.Mod(l2, r2).
func Mod[V1, V2 constraints.Float](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return V1(math.Mod(float64(l1), float64(r1))), V2(math.Mod(float64(l2), float64(r2)))
}

// BinaryAnd returns the binary "and" of l1 and l2 and binary "and" of r1 and r2.
func BinaryAnd[V1, V2 constraints.Integer](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 & r1, l2 & r2
}

// BinaryOr returns the binary "or" of l1 and l2 and binary "or" of r1 and r2.
func BinaryOr[V1, V2 constraints.Integer](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 | r1, l2 | r2
}

// BinaryXor returns the binary "xor" of l1 and l2 and binary "xor" of r1 and r2.
func BinaryXor[V1, V2 constraints.Integer](l1 V1, l2 V2, r1 V1, r2 V2) (V1, V2) {
	return l1 ^ r1, l2 ^ r2
}
