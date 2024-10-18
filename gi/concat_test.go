package gi_test

import (
	"testing"

	"slices"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestConcat(t *testing.T) {
	values1 := slices.Values([]int{1, 2, 3})
	values2 := slices.Values([]int{4, 5, 6})
	values3 := slices.Values([]int{7, 8, 9})

	result := gi.Concat(values1, values2, values3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.Concat(values1)
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.Concat[int]()
	assert.Equal(t, []int(nil), slices.Collect(result))
}

func TestConcatSlices(t *testing.T) {
	values1 := []int{1, 2, 3}
	values2 := []int{4, 5, 6}
	values3 := []int{7, 8, 9}

	result := gi.ConcatSlices(values1, values2, values3)
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.ConcatSlices(values1)
	expected = []int{1, 2, 3}
	assert.Equal(t, expected, slices.Collect(result))

	result = gi.ConcatSlices([]int{})
	assert.Equal(t, []int(nil), slices.Collect(result))
}
