package gi_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/optional/optval"
	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(3), gi.Find(values, gi.Equal(3)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Equal(7)))
}

func TestNotEqual(t *testing.T) {
	values := slices.Values([]int{1})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.NotEqual(4)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.NotEqual(1)))
}

func TestLess(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.Less(4)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Less(1)))
}

func TestLessOrEqual(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.LessOrEqual(4)))
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.LessOrEqual(1)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.LessOrEqual(0)))
}

func TestGreater(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(5), gi.Find(values, gi.Greater(4)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Greater(6)))
}

func TestGreaterOrEqual(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(5), gi.Find(values, gi.GreaterOrEqual(5)))
	assert.Equal(t, optval.Some(6), gi.Find(values, gi.GreaterOrEqual(6)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.GreaterOrEqual(7)))
}

func TestAnd(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(4), gi.Find(values, gi.And(gi.Greater(3), gi.Less(5))))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.And(gi.Greater(3), gi.Less(4))))
}

func TestOr(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(2), gi.Find(values, gi.Or(gi.Greater(3), gi.Even)))
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.Or(gi.Greater(4), gi.Odd)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Or(gi.Greater(6), gi.Less(1))))
}

func TestNot(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.Not(gi.Greater(1))))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Not(gi.Greater(0))))
}

func TestEven(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(2), gi.Find(values, gi.Even))
	values = slices.Values([]int{1, 3, 5})
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Even))
}

func TestOdd(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.Odd))
	values = slices.Values([]int{2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(3), gi.Find(values, gi.Odd))
	values = slices.Values([]int{2, 4, 6})
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.Odd))
}

func TestDivisibleBy(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(3), gi.Find(values, gi.DivisibleBy(3)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.DivisibleBy(7)))
}

func TestIn(t *testing.T) {
	set1 := []int{4, 2, 7}
	set2 := []int{8, 0, 9}
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(2), gi.Find(values, gi.In(set1)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.In(set2)))
}

func TestOneOf(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6})
	assert.Equal(t, optval.Some(3), gi.Find(values, gi.OneOf(5, 4, 3)))
	assert.Equal(t, optval.None[int](), gi.Find(values, gi.OneOf(0, 8, -1)))
}

func TestIsZero(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 0, 4, 5, 6})
	assert.Equal(t, optval.Some(0), gi.Find(values, gi.IsZero))
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.Not[int](gi.IsZero)))
}

func TestIsNotZero(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 0, 4, 5, 6})
	assert.Equal(t, optval.Some(1), gi.Find(values, gi.IsNotZero))
	assert.Equal(t, optval.Some(0), gi.Find(values, gi.Not[int](gi.IsNotZero)))
}
