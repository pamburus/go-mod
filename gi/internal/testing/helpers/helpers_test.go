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
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(2, result)))
}

func TestFlattenPairs(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.FlattenPairs(pairs)

	expected := []int{0, 10, 1, 20, 2, 30}
	assert.Equal(t, expected, slices.Collect(result))

	expected = []int{0, 10, 1, 20}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(4, result)))

	expected = []int{0, 10, 1}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(3, result)))
}

func TestSwap(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.Swap(pairs)

	expected := map[int]int{10: 0, 20: 1, 30: 2}
	assert.Equal(t, expected, maps.Collect(result))

	expected = map[int]int{10: 0, 20: 1}
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(2, result)))
}

func TestLimit(t *testing.T) {
	values := slices.Values([]int{10, 20, 30})

	result := helpers.Limit(2, values)

	expected := []int{10, 20}
	assert.Equal(t, expected, slices.Collect(result))

	expected = []int{10}
	assert.Equal(t, expected, slices.Collect(helpers.Limit(1, result)))
}

func TestLimitPairs(t *testing.T) {
	pairs := slices.All([]int{10, 20, 30})

	result := helpers.LimitPairs(2, pairs)

	expected := map[int]int{0: 10, 1: 20}
	assert.Equal(t, expected, maps.Collect(result))

	expected = map[int]int{0: 10}
	assert.Equal(t, expected, maps.Collect(helpers.LimitPairs(1, result)))
}

func TestNewPair(t *testing.T) {
	assert.Equal(t,
		helpers.Pair[int, int]{42, 43},
		helpers.NewPair(42, 43),
	)
}

func TestPairFold(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6})

	folded := helpers.PairFold(pairs)

	expected := []helpers.Pair[int, int]{
		{V1: 0, V2: 2},
		{V1: 1, V2: 4},
		{V1: 2, V2: 6},
	}
	result := slices.Collect(folded)
	assert.Equal(t, expected, result)

	expected = []helpers.Pair[int, int]{
		{V1: 0, V2: 2},
		{V1: 1, V2: 4},
	}
	result = slices.Collect(helpers.Limit(2, folded))
	assert.Equal(t, expected, result)
}

func TestCollectPairs(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6})

	collected := helpers.CollectPairs(pairs)

	expected := []helpers.Pair[int, int]{
		{V1: 0, V2: 2},
		{V1: 1, V2: 4},
		{V1: 2, V2: 6},
	}
	assert.Equal(t, expected, collected)
}
