package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestSwap(t *testing.T) {
	t.Run("Swap", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4})

		swapped := gi2.Swap(pairs)

		expected := []int{1, 0, 2, 1, 3, 2, 4, 3}
		result := slices.Collect(helpers.FlattenPairs(swapped))
		assert.Equal(t, expected, result)

		expected = []int{1, 0}
		result = slices.Collect(helpers.Limit(2, helpers.FlattenPairs(swapped)))
		assert.Equal(t, expected, result)
	})
}
