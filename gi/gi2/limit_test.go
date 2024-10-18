package gi2_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestLimit(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6, 8})

	limited := gi2.Limit(pairs, 2)
	result := maps.Collect(limited)
	expected := map[int]int{
		0: 2,
		1: 4,
	}

	assert.Equal(t, expected, result)
}

func TestLimitEarlyExit(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6, 8})

	limited := gi2.Limit(pairs, 2)
	result := map[int]int{}

	for k, v := range limited {
		result[k] = v
		if k == 1 {
			break
		}
	}

	expected := map[int]int{
		0: 2,
		1: 4,
	}

	assert.Equal(t, expected, result)
}
