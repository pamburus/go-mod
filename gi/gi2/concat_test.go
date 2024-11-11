package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestConcat(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value + 1
	}

	values1 := slices.Values([]int{1, 3})
	values2 := slices.Values([]int{5, 7})
	values3 := slices.Values([]int{9})

	pairs1 := helpers.ToPairs(values1, makePair)
	pairs2 := helpers.ToPairs(values2, makePair)
	pairs3 := helpers.ToPairs(values3, makePair)

	result := gi2.Concat(pairs1, pairs2, pairs3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, expected, slices.Collect(helpers.FlattenPairs(result)))

	result = gi2.Concat(pairs1)
	assert.Equal(t, []int{1, 2, 3, 4}, slices.Collect(helpers.FlattenPairs(result)))

	result = gi2.Concat[int, int]()
	assert.Equal(t, []int(nil), slices.Collect(helpers.FlattenPairs(result)))
}
