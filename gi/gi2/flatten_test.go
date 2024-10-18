package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFlattenBy(t *testing.T) {
	t.Run("Several", func(t *testing.T) {
		pairs := func(yield func([]int) bool) {
			_ = yield([]int{1, 2}) &&
				yield([]int{3, 4}) &&
				yield([]int{5, 6})
		}

		result := gi2.FlattenBy(pairs, slices.All)

		collected := slices.Collect(helpers.FlattenPairs(result))
		assert.Equal(t, []int{0, 1, 1, 2, 0, 3, 1, 4, 0, 5, 1, 6}, collected)

		collected = slices.Collect(helpers.Limit(helpers.FlattenPairs(result), 3))
		assert.Equal(t, []int{0, 1, 1}, collected)
	})
}
