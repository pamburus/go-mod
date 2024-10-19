package gi2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestEqual(t *testing.T) {
	assert.True(t, gi2.Equal(1, 2)(1, 2))
	assert.False(t, gi2.Equal(1, 2)(1, 3))
}

func TestNotEqual(t *testing.T) {
	assert.True(t, gi2.NotEqual(1, 2)(1, 3))
	assert.False(t, gi2.NotEqual(1, 2)(1, 2))
}

func TestLess(t *testing.T) {
	lt34 := gi2.Less(3, 4)
	assert.True(t, lt34(1, 2))
	assert.True(t, lt34(1, 6))
	assert.False(t, lt34(3, 4))
	assert.False(t, lt34(5, 6))
	assert.False(t, lt34(5, 1))
}

func TestLessOrEqual(t *testing.T) {
	le34 := gi2.LessOrEqual(3, 4)
	assert.True(t, le34(1, 2))
	assert.True(t, le34(1, 6))
	assert.True(t, le34(3, 4))
	assert.False(t, le34(5, 6))
	assert.False(t, le34(5, 1))
}

func TestGreater(t *testing.T) {
	gt34 := gi2.Greater(3, 4)
	assert.False(t, gt34(1, 2))
	assert.False(t, gt34(1, 6))
	assert.False(t, gt34(3, 4))
	assert.True(t, gt34(5, 6))
	assert.True(t, gt34(5, 1))
}

func TestGreaterOrEqual(t *testing.T) {
	ge34 := gi2.GreaterOrEqual(3, 4)
	assert.False(t, ge34(1, 2))
	assert.False(t, ge34(1, 6))
	assert.True(t, ge34(3, 4))
	assert.True(t, ge34(5, 6))
	assert.True(t, ge34(5, 1))
}

func TestAnd(t *testing.T) {
	predicate := gi2.And(
		gi2.Less(3, 4),
		gi2.GreaterOrEqual(1, 2),
	)
	assert.True(t, predicate(1, 3))
	assert.False(t, predicate(1, 1))
	assert.False(t, predicate(3, 4))
	assert.False(t, predicate(5, 6))
	assert.False(t, predicate(5, 1))
}

func TestOr(t *testing.T) {
	predicate := gi2.Or(
		gi2.Less(2, 1),
		gi2.GreaterOrEqual(3, 4),
	)
	assert.True(t, predicate(1, 3))
	assert.True(t, predicate(1, 1))
	assert.False(t, predicate(2, 1))
	assert.False(t, predicate(2, 3))
	assert.False(t, predicate(3, 3))
	assert.True(t, predicate(3, 4))
	assert.True(t, predicate(5, 6))
	assert.True(t, predicate(5, 1))
}

func TestIsZero(t *testing.T) {
	assert.True(t, gi2.IsZero(0, ""))
	assert.False(t, gi2.IsZero(1, ""))
	assert.False(t, gi2.IsZero(0, "zero"))
}

func TestIsNotZero(t *testing.T) {
	assert.False(t, gi2.IsNotZero(0, ""))
	assert.True(t, gi2.IsNotZero(1, ""))
	assert.True(t, gi2.IsNotZero(0, "zero"))
}
