package gi2_test

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name      string
		values    iter.Seq2[int, int]
		expected1 int
		expected2 int
	}{
		{
			"EmptySequence",
			slices.All([]int{}),
			0,
			0,
		},
		{
			"SingleValue",
			slices.All([]int{5}),
			0,
			5,
		},
		{
			"MultipleValues",
			slices.All([]int{1, 2, 3, 4}),
			6,
			10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2 := gi2.Sum(tt.values)
			assert.Equal(t, tt.expected1, v1)
			assert.Equal(t, tt.expected2, v2)
		})
	}
}

func TestProduct(t *testing.T) {
	tests := []struct {
		name      string
		values    iter.Seq2[int, int]
		expected1 int
		expected2 int
	}{
		{
			"EmptySequence",
			slices.All([]int{}),
			1,
			1,
		},
		{
			"SingleValue",
			slices.All([]int{5}),
			0,
			5,
		},
		{
			"MultipleValues",
			slices.All([]int{1, 2, 3, 4}),
			0,
			24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, v2 := gi2.Product(tt.values)
			assert.Equal(t, tt.expected1, v1)
			assert.Equal(t, tt.expected2, v2)
		})
	}
}
