package gi2_test

import (
	"maps"
	"slices"
	"testing"

	"iter"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestMin(t *testing.T) {
	type opt struct {
		v1    int
		v2    int
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq2[int, int]
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.All([]int{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.All([]int{1}),
			expected: opt{v1: 0, v2: 1, valid: true},
		},
		{
			name:     "Multiple",
			values:   helpers.Swap(slices.All([]int{6, 4, 5})),
			expected: opt{v1: 4, v2: 1, valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.Min(tt.values)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMinBy(t *testing.T) {
	type opt struct {
		v1    string
		v2    int
		valid bool
	}

	strLen := func(s string, _ int) int {
		return len(s)
	}

	tests := []struct {
		name     string
		values   iter.Seq2[string, int]
		key      func(string, int) int
		expected opt
	}{
		{
			name:     "Empty",
			values:   helpers.Swap(slices.All([]string{})),
			key:      strLen,
			expected: opt{},
		},
		{
			name:     "Single",
			values:   helpers.Swap(slices.All([]string{"a"})),
			key:      strLen,
			expected: opt{"a", 0, true},
		},
		{
			name:     "Multiple",
			values:   helpers.Swap(slices.All([]string{"abc", "x", "ab"})),
			key:      strLen,
			expected: opt{"x", 1, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MinBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMinByLeft(t *testing.T) {
	type opt struct {
		v1    string
		v2    int
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq2[string, int]
		expected opt
	}{
		{
			name:     "Empty",
			values:   maps.All(map[string]int{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   maps.All(map[string]int{"one": 1}),
			expected: opt{"one", 1, true},
		},
		{
			name:     "Multiple",
			values:   maps.All(map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}),
			expected: opt{"four", 4, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MinByLeft(tt.values)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMinByRight(t *testing.T) {
	type opt struct {
		v1    int
		v2    string
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq2[int, string]
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.All([]string{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.All([]string{"a"}),
			expected: opt{0, "a", true},
		},
		{
			name:     "Multiple",
			values:   slices.All([]string{"abc", "x", "ab"}),
			expected: opt{2, "ab", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MinByRight(tt.values)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}
