package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestFoldPack(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v1, v2 int) int {
			return acc + v1 + v2
		}

		result := gi2.FoldPack(pairs, 0, accumulator)
		assert.Equal(t, 81, result)
	})

	t.Run("SumProduct", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc, v1, v2 int) int {
			return acc + v1*v2
		}

		result := gi2.FoldPack(pairs, 0, accumulator)
		assert.Equal(t, 240, result)
	})
}

func TestFold(t *testing.T) {
	t.Run("Sum", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc1, acc2, v1, v2 int) (int, int) {
			return acc1 + v1, acc2 + v2
		}

		result1, result2 := gi2.Fold(pairs, 0, 0, accumulator)
		assert.Equal(t, 36, result1)
		assert.Equal(t, 45, result2)
	})

	t.Run("SumProduct", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

		accumulator := func(acc1, acc2, v1, v2 int) (int, int) {
			return acc1 + v1, acc2 * v2
		}

		result1, result2 := gi2.Fold(pairs, 0, 1, accumulator)
		assert.Equal(t, 36, result1)
		assert.Equal(t, 362880, result2)
	})
}
