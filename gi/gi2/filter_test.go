package gi2_test

import (
	"testing"

	"slices"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFilter(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value % 3
	}

	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	predicate := func(v1, v2 int) bool {
		return v1%2 == 1 && v2 == 0
	}

	pairs := helpers.ToPairs(values, makePair)

	filtered := gi2.Filter(pairs, predicate)
	result := slices.Collect(helpers.FlattenPairs(filtered))
	expected := []int{3, 0, 9, 0}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.FlattenPairs(helpers.LimitPairs(filtered, 1)))
	expected = []int{3, 0}
	assert.Equal(t, expected, result)
}

func TestFilterLeft(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30, 40, 50, 60})

	predicate := func(v1 int) bool {
		return v1%2 == 0
	}

	filtered := gi2.FilterLeft(pairs, predicate)
	result := slices.Collect(helpers.FlattenPairs(filtered))
	expected := []int{0, 10, 2, 30, 4, 50}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.FlattenPairs(helpers.LimitPairs(filtered, 1)))
	expected = []int{0, 10}
	assert.Equal(t, expected, result)
}

func TestFilterRight(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30, 40, 50, 60})

	predicate := func(v2 int) bool {
		return v2%3 == 0
	}

	filtered := gi2.FilterRight(pairs, predicate)
	result := slices.Collect(helpers.FlattenPairs(filtered))
	expected := []int{2, 30, 5, 60}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.FlattenPairs(helpers.LimitPairs(filtered, 1)))
	expected = []int{2, 30}
	assert.Equal(t, expected, result)
}

func TestFilterKeys(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30, 40, 50, 60})

	predicate := func(index int) bool {
		return index%3 == 0
	}

	filtered := gi2.FilterKeys(pairs, predicate)
	result := slices.Collect(helpers.FlattenPairs(filtered))
	expected := []int{0, 10, 3, 40}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.FlattenPairs(helpers.LimitPairs(filtered, 1)))
	expected = []int{0, 10}
	assert.Equal(t, expected, result)
}

func TestFilterValues(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30, 40, 50, 60})

	predicate := func(value int) bool {
		return value%3 == 0
	}

	filtered := gi2.FilterValues(pairs, predicate)
	result := slices.Collect(helpers.FlattenPairs(filtered))
	expected := []int{2, 30, 5, 60}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.FlattenPairs(helpers.LimitPairs(filtered, 1)))
	expected = []int{2, 30}
	assert.Equal(t, expected, result)
}
