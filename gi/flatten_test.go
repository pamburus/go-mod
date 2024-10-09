package gi_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/stretchr/testify/assert"
)

func TestFlattenBy(t *testing.T) {
	t.Run("Several", func(t *testing.T) {
		values := slices.Values([][]int{{1, 2}, {3, 4}, {5, 6}})

		result := gi.FlattenBy(values, slices.Values)

		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, slices.Collect(result))
		assert.Equal(t, []int{1, 2, 3}, slices.Collect(helpers.Limit(result, 3)))
	})

	t.Run("Single", func(t *testing.T) {
		values := slices.Values([][]int{{1, 2}})

		result := gi.FlattenBy(values, slices.Values)

		assert.Equal(t, []int{1, 2}, slices.Collect(result))
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([][]int{{}})

		result := gi.FlattenBy(values, slices.Values)

		assert.Equal(t, []int(nil), slices.Collect(result))
	})
}

func TestFlattenPairsBy(t *testing.T) {
	t.Run("Several", func(t *testing.T) {
		pairs := func(yield func([]int) bool) {
			_ = yield([]int{1, 2}) &&
				yield([]int{3, 4}) &&
				yield([]int{5, 6})
		}

		result := gi.FlattenPairsBy(pairs, slices.All)

		collected := slices.Collect(helpers.FlattenPairs(result))
		assert.Equal(t, []int{0, 1, 1, 2, 0, 3, 1, 4, 0, 5, 1, 6}, collected)

		collected = slices.Collect(helpers.Limit(helpers.FlattenPairs(result), 3))
		assert.Equal(t, []int{0, 1, 1}, collected)
	})
}
