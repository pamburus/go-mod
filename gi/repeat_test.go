package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestRepeat(t *testing.T) {
	oneToThree := slices.Values([]int{1, 2, 3})

	assert.Equal(t,
		[]int(nil),
		slices.Collect(gi.Repeat(oneToThree, -1)),
	)

	assert.Equal(t,
		[]int(nil),
		slices.Collect(gi.Repeat(oneToThree, 0)),
	)

	assert.Equal(t,
		[]int{1, 2, 3},
		slices.Collect(gi.Repeat(oneToThree, 1)),
	)

	assert.Equal(t,
		[]int{1, 2, 3, 1, 2, 3},
		slices.Collect(gi.Repeat(oneToThree, 2)),
	)

	assert.Equal(t,
		[]int{1, 2, 3, 1, 2, 3, 1},
		slices.Collect(helpers.Limit(7, gi.Repeat(oneToThree, 3))),
	)
}

func TestRepeatSingle(t *testing.T) {
	assert.Equal(t,
		[]int(nil),
		slices.Collect(gi.RepeatSingle(42, -1)),
	)

	assert.Equal(t,
		[]int(nil),
		slices.Collect(gi.RepeatSingle(42, 0)),
	)

	assert.Equal(t,
		[]int{42},
		slices.Collect(gi.RepeatSingle(42, 1)),
	)

	assert.Equal(t,
		[]int{42, 42, 42},
		slices.Collect(helpers.Limit(3, gi.RepeatSingle(42, 5))),
	)
}
