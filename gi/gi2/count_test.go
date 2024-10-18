package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestCount(t *testing.T) {
	makePair := func(value int) (int, int) {
		return value, value%3 + 1
	}

	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	predicate := func(v1, v2 int) bool {
		return v1%2 == 1 && v2%2 == 1
	}

	pairs := helpers.ToPairs(values, makePair)

	result := gi2.Count(pairs)
	assert.Equal(t, 9, result)

	result = gi2.Count(gi2.Filter(pairs, predicate))
	assert.Equal(t, 3, result)
}

func TestCountEmpty(t *testing.T) {
	pairs := slices.All([]int{})

	result := gi2.Count(pairs)
	assert.Equal(t, 0, result)
}
