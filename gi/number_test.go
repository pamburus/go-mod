package gi_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[int]
		expected int
	}{
		{
			"EmptySequence",
			slices.Values([]int{}),
			0,
		},
		{
			"SingleValue",
			slices.Values([]int{5}),
			5,
		},
		{
			"MultipleValues",
			slices.Values([]int{1, 2, 3, 4}),
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.Sum(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		name     string
		values   iter.Seq[int]
		expected int
	}{
		{
			"EmptySequence",
			slices.Values([]int{}),
			1,
		},
		{
			"SingleValue",
			slices.Values([]int{5}),
			5,
		},
		{
			"MultipleValues",
			slices.Values([]int{1, 2, 3, 4}),
			24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := gi.Product(tt.values)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
