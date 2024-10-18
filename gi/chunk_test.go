package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestChunk(t *testing.T) {
	t.Run("Multiple", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("EarlyReturn", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
			if len(chunks) == 2 {
				break
			}
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("RemainderOne", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7},
		}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("RemainderTwo", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8},
		}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("SingleElement", func(t *testing.T) {
		values := slices.Values([]int{1})
		size := 1

		result := gi.Chunk(values, size)
		expected := [][]int{{1}}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("SizeGreaterThanLength", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3})
		size := 5

		result := gi.Chunk(values, size)
		expected := [][]int{{1, 2, 3}}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("ZeroSize", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3})

		assert.Panics(t, func() {
			gi.Chunk(values, 0)
		})
	})

	t.Run("SizeOne", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3})
		size := 1

		result := gi.Chunk(values, size)
		expected := [][]int{{1}, {2}, {3}}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([]int{})
		size := 1

		result := gi.Chunk(values, size)
		var expected [][]int

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, slices.Collect(chunk))
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("PartialReadOne", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1},
			{4, 5},
			{7, 8, 9},
		}

		i := 0
		var chunks [][]int
		for chunk := range result {
			var c []int
			j := 0
			for v := range chunk {
				c = append(c, v)
				j++
				if j == i+1 {
					break
				}
			}
			chunks = append(chunks, c)
			i++
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("PartialReadTwo", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5},
			{7},
		}

		i := 0
		var chunks [][]int
		for chunk := range result {
			var c []int
			j := 0
			for v := range chunk {
				c = append(c, v)
				j++
				if j == 3-i {
					break
				}
			}
			chunks = append(chunks, c)
			i++
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("PartialReadRepeat", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.Chunk(values, size)
		expected := [][]int{
			{1},
			{4, 5},
			{7, 8, 9},
		}

		i := 0
		var chunks [][]int
		for chunk := range result {
			var c []int
			j := 0
			for v := range chunk {
				c = append(c, v)
				j++
				if j == i+1 {
					break
				}
			}
			c = slices.AppendSeq(c, chunk)
			chunks = append(chunks, c)
			i++
		}

		assert.Equal(t, expected, chunks)
	})
}

func TestChunkToSlices(t *testing.T) {
	t.Run("Multiple", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.ChunkToSlices(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		chunks := slices.Collect(result)
		assert.Equal(t, expected, chunks)
	})

	t.Run("EarlyReturn", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
		size := 3

		result := gi.ChunkToSlices(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}

		var chunks [][]int
		for chunk := range result {
			chunks = append(chunks, chunk)
			if len(chunks) == 2 {
				break
			}
		}

		assert.Equal(t, expected, chunks)
	})

	t.Run("RemainderOne", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
		size := 3

		result := gi.ChunkToSlices(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7},
		}

		chunks := slices.Collect(result)
		assert.Equal(t, expected, chunks)
	})

	t.Run("RemainderTwo", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8})
		size := 3

		result := gi.ChunkToSlices(values, size)
		expected := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8},
		}

		chunks := slices.Collect(result)
		assert.Equal(t, expected, chunks)
	})

	t.Run("Panic", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3})

		assert.Panics(t, func() {
			gi.ChunkToSlices(values, 0)
		})
	})
}
