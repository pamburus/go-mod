package gi2_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestChunk(t *testing.T) {
	t.Run("Multiple", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi2.Chunk(pairs, size)
		expected := []map[int]int{
			{0: 1, 1: 2, 2: 3},
			{3: 4, 4: 5, 5: 6},
			{6: 7, 7: 8, 8: 9},
		}

		var chunks []map[int]int
		for chunk := range result {
			chunks = append(chunks, maps.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})
}
