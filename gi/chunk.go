package gi

import (
	"iter"

	"github.com/pamburus/go-mod/gi/constraints"
)

// Chunk returns an iterator that yields chunks of size elements from the given values.
func Chunk[V any, I constraints.Integer](values iter.Seq[V], size I) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("chunk size must be greater than zero")
	}

	return func(yield func(iter.Seq[V]) bool) {
		next, stop := iter.Pull(values)
		defer stop()

		var v V
		hasMore := true
		i := I(1)

		for hasMore {
			if i > 0 {
				// skip remaining values from the current chunk
				// and prefetch the next value
				v, hasMore = next()
				i--

				continue
			}

			i = size
			valid := true
			needMore := yield(func(yield func(V) bool) {
				for ; valid && i > 0; i-- {
					if valid = yield(v); !valid {
						return
					}
					v, hasMore = next()
					if !hasMore {
						return
					}
				}
				valid = false
			})
			if !needMore {
				return
			}
		}
	}
}

// ChunkToSlices returns a new iterator that yields slices of size elements from the given values.
func ChunkToSlices[V any](values iter.Seq[V], size int) iter.Seq[[]V] {
	if size <= 0 {
		panic("chunk size must be greater than zero")
	}

	newChunk := func() []V {
		return make([]V, 0, size)
	}

	return func(yield func([]V) bool) {
		i := 0
		chunk := newChunk()

		for value := range values {
			chunk = append(chunk, value)
			i++
			if i == size {
				if !yield(chunk) {
					return
				}

				chunk = newChunk()
				i = 0
			}
		}

		if len(chunk) > 0 {
			yield(chunk)
		}
	}
}
