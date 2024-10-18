package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestCount(t *testing.T) {
	assert.Equal(t, 4, gi.Count(slices.Values([]int{2, 4, 6, 8})))
	assert.Equal(t, 0, gi.Count(slices.Values([]int{})))
	assert.Equal(t, 1, gi.Count(slices.Values([]int{1})))
}
