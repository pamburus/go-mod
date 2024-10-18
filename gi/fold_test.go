package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestFold(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v int) int {
			return acc + v
		}

		result := gi.Fold(values, 0, accumulator)
		assert.Equal(t, 45, result)
	})

	t.Run("Product", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v int) int {
			return acc * v
		}

		result := gi.Fold(values, 1, accumulator)
		assert.Equal(t, 362880, result)
	})
}
