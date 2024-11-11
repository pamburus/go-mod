package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFilter(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})

	predicate := func(v int) bool {
		return v%2 == 0
	}

	result := slices.Collect(gi.Filter(values, predicate))
	expected := []int{2, 4, 6, 8, 0}
	assert.Equal(t, expected, result)

	result = slices.Collect(helpers.Limit(3, gi.Filter(values, predicate)))
	expected = []int{2, 4, 6}
	assert.Equal(t, expected, result)

	result = slices.Collect(gi.Filter(values, gi.IsNotZero))
	expected = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, result)
}
