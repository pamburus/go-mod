package giop_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giop"
)

func TestAdd(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 5, giop.Add(2, 3))
		assert.Equal(t, -1, giop.Add(2, -3))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 5.0, giop.Add(2.5, 2.5), 1e-9)
		assert.InDelta(t, -1.0, giop.Add(2.5, -3.5), 1e-9)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, -1, giop.Subtract(2, 3))
		assert.Equal(t, 5, giop.Subtract(2, -3))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 0.0, giop.Subtract(2.5, 2.5), 1e-9)
		assert.InDelta(t, 6.0, giop.Subtract(2.5, -3.5), 1e-9)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 6, giop.Multiply(2, 3))
		assert.Equal(t, -6, giop.Multiply(2, -3))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 6.25, giop.Multiply(2.5, 2.5), 1e-9)
		assert.InDelta(t, -8.75, giop.Multiply(2.5, -3.5), 1e-9)
	})
}

func TestDivide(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, 0, giop.Divide(2, 3))
		assert.Equal(t, 2, giop.Divide(6, 3))
		assert.Equal(t, -2, giop.Divide(2, -1))
	})

	t.Run("Float", func(t *testing.T) {
		assert.InDelta(t, 1.0, giop.Divide(2.5, 2.5), 1e-9)
		assert.InDelta(t, -0.625, giop.Divide(2.5, -4), 1e-9)
	})
}

func TestIntMod(t *testing.T) {
	assert.Equal(t, 2, giop.IntMod(2, 3))
	assert.Equal(t, 0, giop.IntMod(6, 3))
	assert.Equal(t, 2, giop.IntMod(2, -3))
}

func TestMod(t *testing.T) {
	assert.InDelta(t, 0.5, giop.Mod(2.5, 2.0), 1e-9)
	assert.InDelta(t, 0.5, giop.Mod(2.5, -2.0), 1e-9)
}

func TestBinaryAnd(t *testing.T) {
	assert.Equal(t, 2, giop.BinaryAnd(2, 3))
	assert.Equal(t, 0, giop.BinaryAnd(2, -3))
}

func TestBinaryOr(t *testing.T) {
	assert.Equal(t, 3, giop.BinaryOr(2, 3))
	assert.Equal(t, -1, giop.BinaryOr(2, -3))
}

func TestBinaryXor(t *testing.T) {
	assert.Equal(t, 1, giop.BinaryXor(2, 3))
	assert.Equal(t, -1, giop.BinaryXor(2, -3))
}
