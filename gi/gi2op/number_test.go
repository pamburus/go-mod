package gi2op_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2op"
)

func TestAdd(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{3, 5}, group(gi2op.Add(1, 2, 2, 3)))
		assert.Equal(t, []any{5, 6}, group(gi2op.Add(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{3.5, 5.5}, group(gi2op.Add(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{7.0, 5.5}, group(gi2op.Add(4.5, 2.0, 2.5, 3.5)))
	})
}

func TestSubtract(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{-1, -1}, group(gi2op.Subtract(1, 2, 2, 3)))
		assert.Equal(t, []any{1, -2}, group(gi2op.Subtract(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{-0.5, -1.5}, group(gi2op.Subtract(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{2.0, -1.5}, group(gi2op.Subtract(4.5, 2.0, 2.5, 3.5)))
	})
}

func TestMultiply(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{2, 6}, group(gi2op.Multiply(1, 2, 2, 3)))
		assert.Equal(t, []any{6, 8}, group(gi2op.Multiply(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{3.0, 7.0}, group(gi2op.Multiply(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{11.25, 8.0}, group(gi2op.Multiply(4.5, 2.0, 2.5, 4)))
	})
}

func TestDivide(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{0, 0}, group(gi2op.Divide(1, 2, 2, 3)))
		assert.Equal(t, []any{1, 0}, group(gi2op.Divide(3, 2, 2, 4)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{0.75, 0.5714285714285714}, group(gi2op.Divide(1.5, 2.0, 2.0, 3.5)))
		assert.Equal(t, []any{1.8, 0.4}, group(gi2op.Divide(4.5, 2.0, 2.5, 5)))
	})
}

func TestIntMod(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	assert.Equal(t, []any{1, 2}, group(gi2op.IntMod(1, 2, 2, 3)))
	assert.Equal(t, []any{1, 3}, group(gi2op.IntMod(3, 7, 2, 4)))
}

func TestMod(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	assert.Equal(t, []any{1.5, 2.5}, group(gi2op.Mod(1.5, 2.5, 2.5, 3.5)))
	assert.Equal(t, []any{1.5, 2.5}, group(gi2op.Mod(1.5, 2.5, 2.5, 3.5)))
}

func TestBinaryAnd(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	assert.Equal(t, []any{0, 2}, group(gi2op.BinaryAnd(1, 2, 2, 3)))
	assert.Equal(t, []any{2, 4}, group(gi2op.BinaryAnd(3, 7, 2, 4)))
}

func TestBinaryOr(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	assert.Equal(t, []any{3, 3}, group(gi2op.BinaryOr(1, 2, 2, 3)))
	assert.Equal(t, []any{3, 7}, group(gi2op.BinaryOr(3, 7, 2, 4)))
}

func TestBinaryXor(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	assert.Equal(t, []any{3, 1}, group(gi2op.BinaryXor(1, 2, 2, 3)))
	assert.Equal(t, []any{1, 3}, group(gi2op.BinaryXor(3, 7, 2, 4)))
}
