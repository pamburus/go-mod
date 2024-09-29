package gi_test

import (
	"slices"
	"testing"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/giop"
	"github.com/stretchr/testify/assert"
)

func TestPairFold(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6})

	expected := []int{2, 5, 8}
	result := slices.Collect(gi.PairFold(pairs, giop.Add))

	assert.Equal(t, expected, result)
}

func TestPairFoldEarlyExit(t *testing.T) {
	pairs := slices.All([]int{2, 4, 6})

	expected := []int{2, 5}
	var result []int

	for v := range gi.PairFold(pairs, giop.Add) {
		result = append(result, v)
		if v == 5 {
			break
		}
	}

	assert.Equal(t, expected, result)
}
