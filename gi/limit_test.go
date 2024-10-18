package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestLimit(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	limited := gi.Limit(values, 3)
	result := slices.Collect(limited)

	expected := []int{1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestLimitEarlyExit(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	limited := gi.Limit(values, 3)
	result := []int{}

	for v := range limited {
		result = append(result, v)
		if v == 2 {
			break
		}
	}

	expected := []int{1, 2}
	assert.Equal(t, expected, result)
}
