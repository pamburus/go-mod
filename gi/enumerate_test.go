package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestEnumerate(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	pairs := gi.Enumerate(values)
	expected := []helpers.Pair[int, int]{
		{V1: 0, V2: 1},
		{V1: 1, V2: 2},
		{V1: 2, V2: 3},
		{V1: 3, V2: 4},
		{V1: 4, V2: 5},
	}
	assert.Equal(t, expected, slices.Collect(helpers.PairFold(pairs)))

	pairs = helpers.LimitPairs(pairs, 3)
	expected = []helpers.Pair[int, int]{
		{V1: 0, V2: 1},
		{V1: 1, V2: 2},
		{V1: 2, V2: 3},
	}
	assert.Equal(t, expected, slices.Collect(helpers.PairFold(pairs)))
}
