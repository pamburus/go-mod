package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestCount(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	predicate := func(v int) bool {
		return v%2 == 0
	}

	result := gi.Count(values, predicate)
	assert.Equal(t, 4, result)
}

func TestCountPairs(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value%3 + 1
	}

	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	predicate := func(v1, v2 int) bool {
		return v1%2 == 1 && v2%2 == 1
	}

	pairs := helpers.ToPairs(values, makePair)

	result := gi.CountPairs(pairs, predicate)
	assert.Equal(t, 3, result)
}

func TestCountPairsEmpty(t *testing.T) {
	pairs := slices.All([]int{})

	predicate := func(int, int) bool {
		return true
	}

	result := gi.CountPairs(pairs, predicate)
	assert.Equal(t, 0, result)
}
