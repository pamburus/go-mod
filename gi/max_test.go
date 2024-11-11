package gi_test

import (
	"cmp"
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestMax(t *testing.T) {
	type opt struct {
		value int
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq[int]
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]int{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]int{1}),
			expected: opt{1, true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]int{3, 1, 2}),
			expected: opt{3, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.Max(tt.values)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}

func TestMaxBy(t *testing.T) {
	type opt struct {
		value string
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq[string]
		key      func(string) int
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			key:      strLen,
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			key:      strLen,
			expected: opt{"a", true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			key:      strLen,
			expected: opt{"abc", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.MaxBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}

func TestMaxByLess(t *testing.T) {
	type opt struct {
		value strLessByLen
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq[strLessByLen]
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]strLessByLen{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]strLessByLen{"a"}),
			expected: opt{strLessByLen("a"), true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]strLessByLen{"abc", "a", "ab"}),
			expected: opt{strLessByLen("abc"), true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.MaxByLess(tt.values)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}

func TestMaxByLessFunc(t *testing.T) {
	type opt struct {
		value string
		valid bool
	}

	less := func(a, b string) bool {
		return len(a) < len(b)
	}

	tests := []struct {
		name     string
		values   iter.Seq[string]
		less     func(string, string) bool
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			less:     less,
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			less:     less,
			expected: opt{"a", true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			less:     less,
			expected: opt{"abc", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.MaxByLessFunc(tt.values, tt.less)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}

func TestMaxByCompare(t *testing.T) {
	type opt struct {
		value strCompareByLen
		valid bool
	}

	tests := []struct {
		name     string
		values   iter.Seq[strCompareByLen]
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]strCompareByLen{}),
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]strCompareByLen{"a"}),
			expected: opt{strCompareByLen("a"), true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]strCompareByLen{"abc", "a", "ab"}),
			expected: opt{strCompareByLen("abc"), true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.MaxByCompare(tt.values)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}

func TestMaxByCompareFunc(t *testing.T) {
	type opt struct {
		value string
		valid bool
	}

	compare := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	tests := []struct {
		name     string
		values   iter.Seq[string]
		compare  func(string, string) int
		expected opt
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			compare:  compare,
			expected: opt{},
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			compare:  compare,
			expected: opt{"a", true},
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			compare:  compare,
			expected: opt{"abc", true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, ok := gi.MaxByCompareFunc(tt.values, tt.compare)
			assert.Equal(t, tt.expected, opt{actual, ok})
		})
	}
}
