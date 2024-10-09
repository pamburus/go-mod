package gi_test

import (
	"testing"

	"slices"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	values1 := slices.Values([]int{1, 2, 3})
	values2 := slices.Values([]int{4, 5, 6})
	values3 := slices.Values([]int{7, 8, 9})

	result := gi.Concat(values1, values2, values3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.Concat(values1)
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.Concat[int]()
	assert.Equal(t, []int(nil), slices.Collect(result))
}

func TestConcatPairs(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value + 1
	}

	values1 := slices.Values([]int{1, 3})
	values2 := slices.Values([]int{5, 7})
	values3 := slices.Values([]int{9})

	pairs1 := helpers.ToPairs(values1, makePair)
	pairs2 := helpers.ToPairs(values2, makePair)
	pairs3 := helpers.ToPairs(values3, makePair)

	result := gi.ConcatPairs(pairs1, pairs2, pairs3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, slices.Collect(helpers.FlattenPairs(result)))

	result = gi.ConcatPairs(pairs1)
	assert.Equal(t, []int{1, 2, 3, 4}, slices.Collect(helpers.FlattenPairs(result)))

	result = gi.ConcatPairs[int, int]()
	assert.Equal(t, []int(nil), slices.Collect(helpers.FlattenPairs(result)))
}

func TestConcatSlices(t *testing.T) {
	values1 := []int{1, 2, 3}
	values2 := []int{4, 5, 6}
	values3 := []int{7, 8, 9}

	result := gi.ConcatSlices(values1, values2, values3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.ConcatSlices(values1)
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.ConcatSlices([]int{})
	assert.Equal(t, []int(nil), slices.Collect(result))
}
