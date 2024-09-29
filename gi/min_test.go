package gi_test

import (
	"cmp"
	"slices"
	"testing"

	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/optional"
	"github.com/stretchr/testify/assert"
)

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[int]
		expected optional.Value[int]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]int{}),
			expected: optional.None[int](),
		},
		{
			name:     "Single",
			values:   slices.Values([]int{1}),
			expected: optional.Some(1),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]int{3, 1, 2}),
			expected: optional.Some(1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.Min(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinBy(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[string]
		key      func(string) int
		expected optional.Value[string]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			key:      strLen,
			expected: optional.None[string](),
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			key:      strLen,
			expected: optional.Some("a"),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			key:      strLen,
			expected: optional.Some("a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.MinBy(tt.values, tt.key)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinByLess(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[strLessByLen]
		expected optional.Value[strLessByLen]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]strLessByLen{}),
			expected: optional.None[strLessByLen](),
		},
		{
			name:     "Single",
			values:   slices.Values([]strLessByLen{"a"}),
			expected: optional.Some(strLessByLen("a")),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]strLessByLen{"abc", "a", "ab"}),
			expected: optional.Some(strLessByLen("a")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.MinByLess(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinByLessFunc(t *testing.T) {
	less := func(a, b string) bool {
		return len(a) < len(b)
	}

	tests := []struct {
		name     string
		values   iter.Seq[string]
		less     func(string, string) bool
		expected optional.Value[string]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			less:     less,
			expected: optional.None[string](),
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			less:     less,
			expected: optional.Some("a"),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			less:     less,
			expected: optional.Some("a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.MinByLessFunc(tt.values, tt.less)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinByCompare(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[strCompareByLen]
		expected optional.Value[strCompareByLen]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]strCompareByLen{}),
			expected: optional.None[strCompareByLen](),
		},
		{
			name:     "Single",
			values:   slices.Values([]strCompareByLen{"a"}),
			expected: optional.Some(strCompareByLen("a")),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]strCompareByLen{"abc", "a", "ab"}),
			expected: optional.Some(strCompareByLen("a")),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.MinByCompare(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestMinByCompareFunc(t *testing.T) {
	compare := func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	}

	tests := []struct {
		name     string
		values   iter.Seq[string]
		compare  func(string, string) int
		expected optional.Value[string]
	}{
		{
			name:     "Empty",
			values:   slices.Values([]string{}),
			compare:  compare,
			expected: optional.None[string](),
		},
		{
			name:     "Single",
			values:   slices.Values([]string{"a"}),
			compare:  compare,
			expected: optional.Some("a"),
		},
		{
			name:     "Multiple",
			values:   slices.Values([]string{"abc", "a", "ab"}),
			compare:  compare,
			expected: optional.Some("a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.MinByCompareFunc(tt.values, tt.compare)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

// ---

type strLessByLen string

func (a strLessByLen) Less(b strLessByLen) bool {
	return len(a) < len(b)
}

// ---

type strCompareByLen string

func (a strCompareByLen) Compare(b strCompareByLen) int {
	return cmp.Compare(len(a), len(b))
}

// ---

func strLen(s string) int {
	return len(s)
}
