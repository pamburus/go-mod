package gi_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	values := slices.Values([]int{1, 2, 3})
	transform := func(v int) int {
		return v * 2
	}

	result := gi.Map(values, transform)
	expected := []int{2, 4, 6}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.Limit(gi.Map(values, transform), 2)
	expected = []int{2, 4}
	assert.Equal(t, expected, slices.Collect(result))
}

func TestMapPairs(t *testing.T) {
	pairs := slices.All([]int{2, 4, -2, 5})
	transform := func(v1, v2 int) (int, int) {
		return v1 * 2, v2 * 2
	}

	result := gi.MapPairs(pairs, transform)
	expected := map[int]int{0: 4, 2: 8, 4: -4, 6: 10}
	assert.Equal(t, expected, maps.Collect(result))

	result = gi.LimitPairs(gi.MapPairs(pairs, transform), 3)
	expected = map[int]int{0: 4, 2: 8, 4: -4}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapLeft(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v1 int) int {
		return v1 * 2
	}

	result := gi.MapLeft(pairs, transform)
	expected := map[int]int{0: 1, 2: 3, 4: 5}
	assert.Equal(t, expected, maps.Collect(result))

	result = gi.LimitPairs(gi.MapLeft(pairs, transform), 2)
	expected = map[int]int{0: 1, 2: 3}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapRight(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v2 int) int {
		return v2 * 2
	}

	result := gi.MapRight(pairs, transform)
	expected := map[int]int{0: 2, 1: 6, 2: 10}
	assert.Equal(t, expected, maps.Collect(result))

	result = gi.LimitPairs(gi.MapRight(pairs, transform), 2)
	expected = map[int]int{0: 2, 1: 6}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapKeys(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v1 int) int {
		return v1 * 2
	}
	result := gi.MapKeys(pairs, transform)
	expected := map[int]int{0: 1, 2: 3, 4: 5}
	assert.Equal(t, expected, maps.Collect(result))
}

func TestMapValues(t *testing.T) {
	pairs := slices.All([]int{1, 3, 5})
	transform := func(v2 int) int {
		return v2 * 2
	}
	result := gi.MapValues(pairs, transform)
	expected := map[int]int{0: 2, 1: 6, 2: 10}
	assert.Equal(t, expected, maps.Collect(result))
}
