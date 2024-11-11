package gi2_test

import (
	"iter"
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestMax(t *testing.T) {
	type opt struct {
		v1    int
		v2    int
		valid bool
	}

	tests := []struct {
		name     string
		pairs    iter.Seq2[int, int]
		expected opt
	}{
		{
			name:     "Empty",
			pairs:    slices.All([]int{}),
			expected: opt{},
		},
		{
			name:     "Single",
			pairs:    slices.All([]int{1}),
			expected: opt{v1: 0, v2: 1, valid: true},
		},
		{
			name:     "Multiple",
			pairs:    slices.All([]int{3, 1, 2}),
			expected: opt{v1: 2, v2: 2, valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.Max(tt.pairs)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMaxBy(t *testing.T) {
	right := func(_ int, b string) string {
		return b
	}

	type opt struct {
		v1    int
		v2    string
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq2[int, string]
		key      func(int, string) string
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.All([]string{}),
			key:      right,
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.All([]string{"a"}),
			key:      right,
			expected: opt{v1: 0, v2: "a", valid: true},
		},
		{
			name:     "Multiple",
			values:   slices.All([]string{"abc", "a", "ab"}),
			key:      right,
			expected: opt{v1: 0, v2: "abc", valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MaxBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMaxByLeft(t *testing.T) {
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
			values:   maps.All(map[int]string{5: "a"}),
			expected: opt{v1: 5, v2: "a", valid: true},
		},
		{
			name:     "Multiple",
			values:   maps.All(map[int]string{3: "abc", 42: "x", 12: "ab"}),
			expected: opt{v1: 42, v2: "x", valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MaxByLeft(tt.values)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}

func TestMaxByRight(t *testing.T) {
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
			values:   maps.All(map[string]int{"a": 5}),
			expected: opt{v1: "a", v2: 5, valid: true},
		},
		{
			name:     "Multiple",
			values:   maps.All(map[string]int{"abc": 3, "x": 42, "ab": 12}),
			expected: opt{v1: "x", v2: 42, valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2, ok := gi2.MaxByRight(tt.values)
			assert.Equal(t, tt.expected, opt{v1, v2, ok})
		})
	}
}
