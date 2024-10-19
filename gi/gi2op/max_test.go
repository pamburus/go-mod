package gi2op_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2op"
)

func TestMax(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{2, 3}, group(gi2op.Max(2, 3, 1, 2)))
		assert.Equal(t, []any{2, -3}, group(gi2op.Max(2, -3, 1, 2)))
		assert.Equal(t, []any{2, 3}, group(gi2op.Max(2, 3, 2, 2)))
		assert.Equal(t, []any{2, 2}, group(gi2op.Max(2, -3, 2, 2)))
		assert.Equal(t, []any{3, 3}, group(gi2op.Max(3, 3, 3, 3)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{2.5, 2.5}, group(gi2op.Max(2.5, 2.5, 1.5, 2.5)))
		assert.Equal(t, []any{2.5, -3.5}, group(gi2op.Max(2.5, -3.5, 1.5, 2.5)))
		assert.Equal(t, []any{2.5, 3.5}, group(gi2op.Max(1.5, 3.5, 2.5, 3.5)))
	})
}

func TestMaxBy(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		maxMod2 := func(a, b int) int {
			return max(a%2, b%2)
		}

		assert.Equal(t, []any{2, 3}, group(gi2op.MaxBy(maxMod2)(2, 3, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(gi2op.MaxBy(maxMod2)(3, 4, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(gi2op.MaxBy(maxMod2)(1, 2, 3, 4)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxBy(maxMod2)(2, 2, 3, 3)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxBy(maxMod2)(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		mod2 := func(a, b float64) float64 {
			return max(math.Mod(a, 2), math.Mod(b, 2))
		}

		assert.Equal(t, []any{2.0, 3.0}, group(gi2op.MaxBy(mod2)(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{3.0, 4.0}, group(gi2op.MaxBy(mod2)(3.0, 4.0, 1.0, 2.0)))
	})
}

func TestMaxByLeft(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{2, 3}, group(gi2op.MaxByLeft(2, 3, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(gi2op.MaxByLeft(3, 4, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(gi2op.MaxByLeft(1, 2, 3, 4)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxByLeft(2, 2, 3, 3)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxByLeft(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{2.0, 3.0}, group(gi2op.MaxByLeft(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{3.0, 4.0}, group(gi2op.MaxByLeft(3.0, 4.0, 1.0, 2.0)))
	})
}

func TestMaxByRight(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{2, 3}, group(gi2op.MaxByRight(2, 3, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(gi2op.MaxByRight(3, 4, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(gi2op.MaxByRight(1, 2, 3, 4)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxByRight(2, 2, 3, 3)))
		assert.Equal(t, []any{3, 3}, group(gi2op.MaxByRight(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{2.0, 3.0}, group(gi2op.MaxByRight(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{3.0, 4.0}, group(gi2op.MaxByRight(3.0, 4.0, 1.0, 2.0)))
	})
}
