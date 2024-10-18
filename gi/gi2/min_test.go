package gi2_test

import (
	"slices"
	"testing"

	"iter"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq2[int, int]
		expected optpair.Pair[int, int]
	}{
		{
			name:     "Empty",
			values:   slices.All([]int{}),
			expected: optpair.None[int, int](),
		},
		{
			name:     "Single",
			values:   slices.All([]int{1}),
			expected: optpair.Some(0, 1),
		},
		{
			name:     "Multiple",
			values:   helpers.Swap(slices.All([]int{6, 4, 5})),
			expected: optpair.Some(4, 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi2.Min(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinBy(t *testing.T) {
	strLen := func(s string, _ int) int {
		return len(s)
	}

	tests := []struct {
		name     string
		values   iter.Seq2[string, int]
		key      func(string, int) int
		expected optpair.Pair[string, int]
	}{
		{
			name:     "Empty",
			values:   helpers.Swap(slices.All([]string{})),
			key:      strLen,
			expected: optpair.None[string, int](),
		},
		{
			name:     "Single",
			values:   helpers.Swap(slices.All([]string{"a"})),
			key:      strLen,
			expected: optpair.Some("a", 0),
		},
		{
			name:     "Multiple",
			values:   helpers.Swap(slices.All([]string{"abc", "x", "ab"})),
			key:      strLen,
			expected: optpair.Some("x", 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi2.MinBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
