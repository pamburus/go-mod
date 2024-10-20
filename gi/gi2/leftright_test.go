package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestLeft(t *testing.T) {
	pairs := slices.All([]int{31, 32, 33, 34})

	values := slices.Collect(gi2.Left(pairs))
	assert.Equal(t, values, []int{0, 1, 2, 3})

	values = slices.Collect(helpers.Limit(2, gi2.Left(pairs)))
	assert.Equal(t, values, []int{0, 1})
}

func TestRight(t *testing.T) {
	pairs := slices.All([]int{31, 32, 33, 34})

	values := slices.Collect(gi2.Right(pairs))
	assert.Equal(t, values, []int{31, 32, 33, 34})

	values = slices.Collect(helpers.Limit(2, gi2.Right(pairs)))
	assert.Equal(t, values, []int{31, 32})
}
