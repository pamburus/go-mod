package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
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
