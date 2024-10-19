package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestLoop(t *testing.T) {
	values := slices.Values([]int{1, 2, 3})

	assert.Equal(t,
		[]int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1},
		slices.Collect(helpers.Limit(10, gi.Loop(values))),
	)
}

func TestLoopSingle(t *testing.T) {
	assert.Equal(t,
		[]int{42, 42, 42, 42, 42},
		slices.Collect(helpers.Limit(5, gi.LoopSingle(42))),
	)
}
