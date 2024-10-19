package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/giop"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestPack(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		pairs := slices.All([]int{2, 4, 6})

		packed := gi2.Pack(pairs, giop.Add)

		expected := []int{2, 5, 8}
		result := slices.Collect(packed)
		assert.Equal(t, expected, result)

		expected = []int{2, 5}
		result = slices.Collect(helpers.Limit(2, packed))
		assert.Equal(t, expected, result)
	})
}

func TestUnpack(t *testing.T) {
	divMod2 := func(v int) (int, int) {
		return v / 2, v % 2
	}

	t.Run("DivMod", func(t *testing.T) {
		values := slices.Values([]int{2, 3, 4})

		unpacked := gi2.Unpack(values, divMod2)

		expected := []int{1, 0, 1, 1, 2, 0}
		result := slices.Collect(helpers.FlattenPairs(unpacked))
		assert.Equal(t, expected, result)

		expected = []int{1, 0, 1, 1}
		result = slices.Collect(helpers.Limit(4, helpers.FlattenPairs(unpacked)))
		assert.Equal(t, expected, result)
	})
}
