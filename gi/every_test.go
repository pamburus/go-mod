package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestEvery(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10, 11})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.False(t, result)
	})

	t.Run("AllTrue", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.True(t, result)
	})

	t.Run("AllFalse", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10, 11})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([]int{})

		predicate := func(int) bool {
			return false
		}

		result := gi.Every(values, predicate)
		assert.True(t, result)
	})
}

func TestEveryPair(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value * 2
	}

	t.Run("SomeTrue", func(t *testing.T) {
		values := slices.Values([]int{1, 3, 5, 7, 9})

		predicate := func(v1, v2 int) bool {
			return v2-v1 > 5
		}

		pairs := helpers.ToPairs(values, makePair)

		result := gi.EveryPair(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("AllTrue", func(t *testing.T) {
		values := slices.Values([]int{1, 3, 5, 7, 9})

		predicate := func(v1, v2 int) bool {
			return v1%2 == 1 && v2%2 == 0
		}

		pairs := helpers.ToPairs(values, makePair)

		result := gi.EveryPair(pairs, predicate)
		assert.True(t, result)
	})

	t.Run("AllFalse", func(t *testing.T) {
		values := slices.Values([]int{1, 3, 5, 7, 9})

		predicate := func(v1, v2 int) bool {
			return v1%2 == 0 && v2%2 == 1
		}

		pairs := helpers.ToPairs(values, makePair)

		result := gi.EveryPair(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([]int{})

		pairs := helpers.ToPairs(values, makePair)

		predicate := func(int, int) bool {
			return false
		}

		result := gi.EveryPair(pairs, predicate)
		assert.True(t, result)
	})
}
