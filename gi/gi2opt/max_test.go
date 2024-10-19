package gi2opt_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2opt"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestMax(t *testing.T) {
	assert.Equal(t,
		optpair.New(9, 7, true),
		gi2opt.Max(maps.All(map[int]int{8: 5, 3: 3, 4: 6, 9: 7, 6: 4})),
	)
}

func TestMaxBy(t *testing.T) {
	rightAbs := func(_, r int) int {
		if r < 0 {
			return -r
		}

		return r
	}

	assert.Equal(t,
		optpair.New(3, -7, true),
		gi2opt.MaxBy(slices.All([]int{5, -3, 6, -7, 4}), rightAbs),
	)
}

func TestMaxByLeft(t *testing.T) {
	assert.Equal(t,
		optpair.New(9, 7, true),
		gi2opt.MaxByLeft(maps.All(map[int]int{8: 5, 3: 3, 4: 6, 9: 7, 6: 4})),
	)
}

func TestMaxByRight(t *testing.T) {
	assert.Equal(t,
		optpair.New(2, 6, true),
		gi2opt.MaxByRight(slices.All([]int{5, -3, 6, -7, 4})),
	)
}
