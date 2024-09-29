package gi_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	limited := gi.Limit(values, 3)
	result := slices.Collect(limited)

	expected := []int{1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestLimitEarlyExit(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	limited := gi.Limit(values, 3)
	result := []int{}

	for v := range limited {
		result = append(result, v)
		if v == 2 {
			break
		}
	}

	expected := []int{1, 2}
	assert.Equal(t, expected, result)
}

func TestLimitPairs(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6, 8})

	limited := gi.LimitPairs(pairs, 2)
	result := maps.Collect(limited)
	expected := map[int]int{
		0: 2,
		1: 4,
	}

	assert.Equal(t, expected, result)
}

func TestLimitPairsEarlyExit(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6, 8})

	limited := gi.LimitPairs(pairs, 2)
	result := map[int]int{}

	for k, v := range limited {
		result[k] = v
		if k == 1 {
			break
		}
	}

	expected := map[int]int{
		0: 2,
		1: 4,
	}

	assert.Equal(t, expected, result)
}
