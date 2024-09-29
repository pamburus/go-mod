package gi

import "iter"

// Chunk returns an iterator that yields chunks of size elements from the given values.
func Chunk[V any](values iter.Seq[V], size int) iter.Seq[iter.Seq[V]] {
	if size <= 0 {
		panic("chunk size must be greater than zero")
	}

	return func(yield func(iter.Seq[V]) bool) {
		next, stop := iter.Pull(values)
		defer stop()

		active := true
		i := 0

		for active {
			if i > 0 && i < size {
				_, active = next()
				i++

				continue
			}

			active = yield(func(yield func(V) bool) {
				for {
					if v, ok := next(); ok {
						i++
						if !yield(v) || i == size {
							return
						}
					} else {
						active = false

						return
					}
				}
			})
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
