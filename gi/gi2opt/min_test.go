package gi2opt_test

import (
	"maps"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2opt"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestMin(t *testing.T) {
	assert.Equal(t,
		optpair.New(3, 3, true),
		gi2opt.Min(maps.All(map[int]int{8: 5, 3: 3, 4: 6, 9: 7, 6: 4})),
	)
}

func TestMinBy(t *testing.T) {
	rightAbs := func(_, r int) int {
		if r < 0 {
			return -r
		}

		return r
	}

	assert.Equal(t,
		optpair.New(1, -3, true),
		gi2opt.MinBy(slices.All([]int{5, -3, 6, -7, 4}), rightAbs),
	)
}

func TestMinByLeft(t *testing.T) {
	assert.Equal(t,
		optpair.New(3, 4, true),
		gi2opt.MinByLeft(maps.All(map[int]int{8: 5, 3: 4, 4: 6, 9: 7, 6: 4})),
	)
}

func TestMinByRight(t *testing.T) {
	assert.Equal(t,
		optpair.New(3, -7, true),
		gi2opt.MinByRight(slices.All([]int{5, -3, 6, -7, 4})),
	)
}
