package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFlattenBy(t *testing.T) {
	t.Run("Several", func(t *testing.T) {
		values := slices.Values([][]int{{1, 2}, {3, 4}, {5, 6}})

		result := gi.FlattenBy(values, slices.Values)

		assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, slices.Collect(result))
		assert.Equal(t, []int{1, 2, 3}, slices.Collect(helpers.Limit(3, result)))
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
