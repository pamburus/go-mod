package gi_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	predicate := func(v int) bool {
		return v%2 == 0
	}

	result := gi.Count(values, predicate)
	assert.Equal(t, 4, result)
}
