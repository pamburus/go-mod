package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestReduce(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v int) int {
			return acc + v
		}

		result := gi.Reduce(values, accumulator)
		assert.Equal(t, optval.Some(45), result)
	})

	t.Run("Product", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v int) int {
			return acc * v
		}

		result := gi.Reduce(values, accumulator)
		assert.Equal(t, optval.Some(362880), result)
	})
}
