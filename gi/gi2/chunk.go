package gi2

import (
	"iter"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/constraints"
)

// Chunk returns an iterator that yields chunks of pairs of given size.
func Chunk[V1, V2 any, I constraints.Integer](pairs iter.Seq2[V1, V2], size I) iter.Seq[iter.Seq2[V1, V2]] {
	type pair struct {
		v1 V1
		v2 V2
	}

	pack := PackWith(func(v1 V1, v2 V2) pair {
		return pair{v1, v2}
	})

	unpack := UnpackWith(func(p pair) (V1, V2) {
		return p.v1, p.v2
	})

	return gi.Map(gi.Chunk(pack(pairs), size), unpack)
}
