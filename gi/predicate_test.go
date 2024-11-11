package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestEqual(t *testing.T) {
	assert.True(t, gi.Equal(3)(3))
	assert.False(t, gi.Equal(3)(2))
}

func TestNotEqual(t *testing.T) {
	assert.True(t, gi.NotEqual(3)(2))
	assert.False(t, gi.NotEqual(3)(3))
}

func TestLess(t *testing.T) {
	lt4 := gi.Less(4)
	assert.True(t, lt4(3))
	assert.False(t, lt4(4))
}

func TestLessOrEqual(t *testing.T) {
	le4 := gi.LessOrEqual(4)
	assert.True(t, le4(3))
	assert.True(t, le4(4))
	assert.False(t, le4(5))
}

func TestGreater(t *testing.T) {
	gt4 := gi.Greater(4)
	assert.True(t, gt4(5))
	assert.False(t, gt4(4))
}

func TestGreaterOrEqual(t *testing.T) {
	ge4 := gi.GreaterOrEqual(4)
	assert.True(t, ge4(5))
	assert.True(t, ge4(4))
	assert.False(t, ge4(3))
}

func TestAnd(t *testing.T) {
	lt4 := gi.Less(4)
	gt2 := gi.Greater(2)
	assert.True(t, gi.And(lt4, gt2)(3))
	assert.False(t, gi.And(lt4, gt2)(2))
	assert.False(t, gi.And(lt4, gt2)(4))
}

func TestOr(t *testing.T) {
	lt2 := gi.Less(2)
	gt4 := gi.Greater(4)
	assert.False(t, gi.Or(lt2, gt4)(3))
	assert.False(t, gi.Or(lt2, gt4)(2))
	assert.True(t, gi.Or(lt2, gt4)(1))
	assert.True(t, gi.Or(lt2, gt4)(5))
}

func TestNot(t *testing.T) {
	lt2 := gi.Less(2)
	assert.True(t, gi.Not(lt2)(3))
	assert.False(t, gi.Not(lt2)(1))
}

func TestEven(t *testing.T) {
	assert.True(t, gi.Even(2))
	assert.False(t, gi.Even(3))
}

func TestOdd(t *testing.T) {
	assert.True(t, gi.Odd(3))
	assert.False(t, gi.Odd(2))
}

func TestDivisibleBy(t *testing.T) {
	div2 := gi.DivisibleBy(2)
	assert.True(t, div2(4))
	assert.False(t, div2(3))
}

func TestIn(t *testing.T) {
	set1 := []int{4, 2, 7}
	set2 := []int{8, 0, 9}
	inSet1 := gi.In(set1)
	inSet2 := gi.In(set2)

	assert.True(t, inSet1(4))
	assert.False(t, inSet1(3))
	assert.True(t, inSet2(0))
	assert.False(t, inSet2(3))
}

func TestOneOf(t *testing.T) {
	oneOf123 := gi.OneOf(1, 2, 3)
	assert.False(t, oneOf123(0))
	assert.True(t, oneOf123(1))
	assert.True(t, oneOf123(2))
	assert.True(t, oneOf123(3))
	assert.False(t, oneOf123(4))
}

func TestIsZero(t *testing.T) {
	assert.True(t, gi.IsZero(0))
	assert.False(t, gi.IsZero(1))
	assert.True(t, gi.IsZero(0.0))
	assert.False(t, gi.IsZero(0.1))
	assert.True(t, gi.IsZero(struct{}{}))
	assert.True(t, gi.IsZero(struct{ int }{}))
	assert.False(t, gi.IsZero(struct{ int }{3}))
}

func TestIsNotZero(t *testing.T) {
	assert.False(t, gi.IsNotZero(0))
	assert.True(t, gi.IsNotZero(1))
	assert.False(t, gi.IsNotZero(0.0))
	assert.True(t, gi.IsNotZero(0.1))
	assert.False(t, gi.IsNotZero(struct{}{}))
	assert.False(t, gi.IsNotZero(struct{ int }{}))
	assert.True(t, gi.IsNotZero(struct{ int }{3}))
}

func TestEach(t *testing.T) {
	eachIsEven := gi.Each(gi.Even[int])
	assert.True(t, eachIsEven(slices.Values([]int{2, 4, 6})))
	assert.False(t, eachIsEven(slices.Values([]int{2, 3, 4})))
}

func TestAny(t *testing.T) {
	anyIsEven := gi.Any(gi.Even[int])
	assert.True(t, anyIsEven(slices.Values([]int{2, 3, 4})))
	assert.False(t, anyIsEven(slices.Values([]int{3, 5, 7})))
}
