package giop2_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giop2"
)

func TestMin(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{1, 2}, group(giop2.Min(2, 3, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(giop2.Min(2, -3, 1, 2)))
		assert.Equal(t, []any{2, 2}, group(giop2.Min(2, 3, 2, 2)))
		assert.Equal(t, []any{2, -3}, group(giop2.Min(2, -3, 2, 2)))
		assert.Equal(t, []any{3, 3}, group(giop2.Min(3, 3, 3, 3)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{1.5, 2.5}, group(giop2.Min(2.5, 2.5, 1.5, 2.5)))
		assert.Equal(t, []any{1.5, 2.5}, group(giop2.Min(2.5, -3.5, 1.5, 2.5)))
		assert.Equal(t, []any{1.5, 3.5}, group(giop2.Min(1.5, 3.5, 2.5, 3.5)))
	})
}

func TestMinBy(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		minMod2 := func(a, b int) int {
			return min(a%2, b%2)
		}

		assert.Equal(t, []any{2, 3}, group(giop2.MinBy(minMod2)(2, 3, 1, 2)))
		assert.Equal(t, []any{3, 4}, group(giop2.MinBy(minMod2)(3, 4, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(giop2.MinBy(minMod2)(1, 2, 3, 4)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinBy(minMod2)(2, 2, 3, 3)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinBy(minMod2)(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		mod2 := func(a, b float64) float64 {
			return min(math.Mod(a, 2), math.Mod(b, 2))
		}

		assert.Equal(t, []any{2.0, 3.0}, group(giop2.MinBy(mod2)(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{3.0, 4.0}, group(giop2.MinBy(mod2)(3.0, 4.0, 1.0, 2.0)))
	})
}

func TestMinByLeft(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{1, 2}, group(giop2.MinByLeft(2, 3, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(giop2.MinByLeft(3, 4, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(giop2.MinByLeft(1, 2, 3, 4)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinByLeft(2, 2, 3, 3)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinByLeft(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{1.0, 2.0}, group(giop2.MinByLeft(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{1.0, 2.0}, group(giop2.MinByLeft(3.0, 4.0, 1.0, 2.0)))
	})
}

func TestMinByRight(t *testing.T) {
	group := func(values ...any) []any {
		return values
	}

	t.Run("Int", func(t *testing.T) {
		assert.Equal(t, []any{1, 2}, group(giop2.MinByRight(2, 3, 1, 2)))
		assert.Equal(t, []any{3, 1}, group(giop2.MinByRight(3, 1, 1, 2)))
		assert.Equal(t, []any{1, 2}, group(giop2.MinByRight(1, 2, 3, 4)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinByRight(2, 2, 3, 3)))
		assert.Equal(t, []any{2, 2}, group(giop2.MinByRight(3, 3, 2, 2)))
	})

	t.Run("Float", func(t *testing.T) {
		assert.Equal(t, []any{1.0, 2.0}, group(giop2.MinByRight(2.0, 3.0, 1.0, 2.0)))
		assert.Equal(t, []any{3.0, 4.0}, group(giop2.MinByRight(3.0, 4.0, 1.0, 9.0)))
	})
}