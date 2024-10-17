package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestLoop(t *testing.T) {
	pairs := slices.All([]int{1, 2, 3})

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 0, V2: 1},
			{V1: 1, V2: 2},
			{V1: 2, V2: 3},
			{V1: 0, V2: 1},
			{V1: 1, V2: 2},
		},
		helpers.CollectPairs(helpers.LimitPairs(5, gi2.Loop(pairs))),
	)
}

func TestLoopSingle(t *testing.T) {
	assert.Equal(t,
		[]helpers.Pair[int, string]{
			{V1: 42, V2: "foo"},
			{V1: 42, V2: "foo"},
			{V1: 42, V2: "foo"},
		},
		helpers.CollectPairs(helpers.LimitPairs(3, gi2.LoopSingle(42, "foo"))),
	)
}
