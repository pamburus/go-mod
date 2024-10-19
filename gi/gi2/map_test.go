package gi2_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestMap(t *testing.T) {
	pairs := slices.All([]int{2, 4, -2, 5})
	transform := func(v1, v2 int) (int, int) {
		return v1 * 2, v2 * 2
	}

	result := gi2.Map(pairs, transform)
	expected := map[int]int{0: 4, 2: 8, 4: -4, 6: 10}
	assert.Equal(t, expected, maps.Collect(result))

	result = helpers.LimitPairs(3, gi2.Map(pairs, transform))
	expected = map[int]int{0: 4, 2: 8, 4: -4}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapLeft(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v1 int) int {
		return v1 * 2
	}

	result := gi2.MapLeft(pairs, transform)
	expected := map[int]int{0: 1, 2: 3, 4: 5}
	assert.Equal(t, expected, maps.Collect(result))

	result = helpers.LimitPairs(2, gi2.MapLeft(pairs, transform))
	expected = map[int]int{0: 1, 2: 3}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapRight(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v2 int) int {
		return v2 * 2
	}

	result := gi2.MapRight(pairs, transform)
	expected := map[int]int{0: 2, 1: 6, 2: 10}
	assert.Equal(t, expected, maps.Collect(result))

	result = helpers.LimitPairs(2, gi2.MapRight(pairs, transform))
	expected = map[int]int{0: 2, 1: 6}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapKeys(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v1 int) int {
		return v1 * 2
	}
	result := gi2.MapKeys(pairs, transform)
	expected := map[int]int{0: 1, 2: 3, 4: 5}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapValues(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v2 int) int {
		return v2 * 2
	}
	result := gi2.MapValues(pairs, transform)
	expected := map[int]int{0: 2, 1: 6, 2: 10}
	assert.Equal(t, expected, maps.Collect(result))
}
