package gi2opt_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2opt"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestFilterMap(t *testing.T) {
	assert.Equal(t,
		[]helpers.Pair[int, int]{
			{V1: 0, V2: 50},
			{V1: 2, V2: 60},
			{V1: 4, V2: 40},
		},
		helpers.CollectPairs(
			gi2opt.FilterMap(
				slices.All([]int{5, 3, 6, 7, 4}),
				func(i, v int) optpair.Pair[int, int] {
					return optpair.New(i, v*10, i%2 == 0)
				},
			),
		),
	)
}
