package gi2_test

import (
	"slices"
	"testing"

	"iter"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		pairs    iter.Seq2[int, int]
		expected optpair.Pair[int, int]
	}{
		{
			name:     "Empty",
			pairs:    slices.All([]int{}),
			expected: optpair.None[int, int](),
		},
		{
			name:     "Single",
			pairs:    slices.All([]int{1}),
			expected: optpair.Some(0, 1),
		},
		{
			name:     "Multiple",
			pairs:    slices.All([]int{3, 1, 2}),
			expected: optpair.Some(2, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi2.Max(tt.pairs)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMaxBy(t *testing.T) {
	right := func(_ int, b string) string {
		return b
	}

	tests := []struct {
		name     string
		values   iter.Seq2[int, string]
		key      func(int, string) string
		expected optpair.Pair[int, string]
	}{
		{
			name:     "Empty",
			values:   slices.All([]string{}),
			key:      right,
			expected: optpair.None[int, string](),
		},
		{
			name:     "Single",
			values:   slices.All([]string{"a"}),
			key:      right,
			expected: optpair.Some(0, "a"),
		},
		{
			name:     "Multiple",
			values:   slices.All([]string{"abc", "a", "ab"}),
			key:      right,
			expected: optpair.Some(0, "abc"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi2.MaxBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
