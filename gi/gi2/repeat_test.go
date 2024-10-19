package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestRepeat(t *testing.T) {
	values := slices.All([]int{7, 42, 123})

	assert.Equal(t,
		[]helpers.Pair[int, int](nil),
		helpers.CollectPairs(gi2.Repeat(values, -1)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int](nil),
		helpers.CollectPairs(gi2.Repeat(values, 0)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 0, V2: 7},
			{V1: 1, V2: 42},
			{V1: 2, V2: 123},
		},
		helpers.CollectPairs(gi2.Repeat(values, 1)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 0, V2: 7},
			{V1: 1, V2: 42},
			{V1: 2, V2: 123},
			{V1: 0, V2: 7},
			{V1: 1, V2: 42},
			{V1: 2, V2: 123},
		},
		helpers.CollectPairs(gi2.Repeat(values, 2)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 0, V2: 7},
			{V1: 1, V2: 42},
			{V1: 2, V2: 123},
			{V1: 0, V2: 7},
			{V1: 1, V2: 42},
			{V1: 2, V2: 123},
			{V1: 0, V2: 7},
		},
		helpers.CollectPairs(helpers.LimitPairs(7, gi2.Repeat(values, 3))),
	)
}

func TestRepeatSingle(t *testing.T) {
	assert.Equal(t,
		[]helpers.Pair[int, int](nil),
		helpers.CollectPairs(gi2.RepeatSingle(42, 7, -1)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int](nil),
		helpers.CollectPairs(gi2.RepeatSingle(1, 42, 0)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 123, V2: 42},
		},
		helpers.CollectPairs(gi2.RepeatSingle(123, 42, 1)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 123, V2: 42},
			{V1: 123, V2: 42},
			{V1: 123, V2: 42},
		},
		helpers.CollectPairs(helpers.LimitPairs(3, gi2.RepeatSingle(123, 42, 5))),
	)
}
