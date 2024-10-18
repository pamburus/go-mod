package helpers_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestToPairs(t *testing.T) {
	values := []int{1, 2, 3}

	toPair := func(v int) (int, int) {
		return v, v
	}

	result := helpers.ToPairs(slices.Values(values), toPair)

	expected := map[int]int{1: 1, 2: 2, 3: 3}
	assert.Equal(t, expected, maps.Collect(result))

	expected = map[int]int{1: 1, 2: 2}
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(result, 2)))
}

func TestFlattenPairs(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.FlattenPairs(pairs)

	expected := []int{0, 10, 1, 20, 2, 30}
	assert.Equal(t, expected, slices.Collect(result))

	expected = []int{0, 10, 1, 20}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(result, 4)))

	expected = []int{0, 10, 1}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(result, 3)))
}

func TestSwap(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.Swap(pairs)

	expected := map[int]int{10: 0, 20: 1, 30: 2}
	assert.Equal(t, expected, maps.Collect(result))

	expected = map[int]int{10: 0, 20: 1}
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(result, 2)))
}

func TestLimit(t *testing.T) {
	values := slices.Values([]int{10, 20, 30})

	result := helpers.Limit(values, 2)

	expected := []int{10, 20}
	assert.Equal(t, expected, slices.Collect(result))

	expected = []int{10}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(result, 1)))
}

func TestLimitPairs(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.LimitPairs(pairs, 2)

	expected := map[int]int{0: 10, 1: 20}
	assert.Equal(t, expected, maps.Collect(result))

	expected = map[int]int{0: 10}
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(result, 1)))
}
