package gi_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/giop"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPairFold(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		pairs := slices.All([]int{2, 4, 6})

		folded := gi.PairFold(pairs, giop.Add)

		expected := []int{2, 5, 8}
		result := slices.Collect(folded)
		assert.Equal(t, expected, result)

		expected = []int{2, 5}
		result = slices.Collect(helpers.Limit(folded, 2))
		assert.Equal(t, expected, result)
	})
}

func TestPairUnfold(t *testing.T) {
	div := func(v int) (int, int) {
		return v / 2, v % 2
	}

	t.Run("DivMod", func(t *testing.T) {
		values := slices.Values([]int{2, 3, 4})

		unfolded := gi.PairUnfold(values, div)

		expected := []int{1, 0, 1, 1, 2, 0}
		result := slices.Collect(helpers.FlattenPairs(unfolded))
		assert.Equal(t, expected, result)

		expected = []int{1, 0, 1, 1}
		result = slices.Collect(helpers.Limit(helpers.FlattenPairs(unfolded), 4))
		assert.Equal(t, expected, result)
	})
}

func TestPairSwap(t *testing.T) {
	t.Run("Swap", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4})

		swapped := gi.PairSwap(pairs)

		expected := []int{1, 0, 2, 1, 3, 2, 4, 3}
		result := slices.Collect(helpers.FlattenPairs(swapped))
		assert.Equal(t, expected, result)

		expected = []int{1, 0}
		result = slices.Collect(helpers.Limit(helpers.FlattenPairs(swapped), 2))
		assert.Equal(t, expected, result)
	})
}
