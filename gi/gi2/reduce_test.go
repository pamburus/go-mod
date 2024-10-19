package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestReduce(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc1, acc2, v1, v2 int) (int, int) {
			return acc1 + v1, acc2 + v2
		}

		v1, v2, ok := gi2.Reduce(pairs, accumulator)
		require.True(t, ok)
		assert.Equal(t, 36, v1)
		assert.Equal(t, 45, v2)
	})

	t.Run("Product", func(t *testing.T) {
		values := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc1, acc2, v1, v2 int) (int, int) {
			return acc1 * v1, acc2 * v2
		}

		v1, v2, ok := gi2.Reduce(values, accumulator)
		require.True(t, ok)
		assert.Equal(t, 0, v1)
		assert.Equal(t, 362880, v2)
	})
}
