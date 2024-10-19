package gi2_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFilterMap(t *testing.T) {
	empty := slices.All([]string{})
	zeroToFour := slices.All([]string{"zero", "one", "two", "three", "four"})

	evenNeg := func(i int, v string) (int, string, bool) {
		if i%2 == 0 {
			return -i, v, true
		}

		return 0, "", false
	}

	assert.Equal(t,
		[]helpers.Pair[int, string](nil),
		helpers.CollectPairs(gi2.FilterMap(empty, evenNeg)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, string]{
			{V1: 0, V2: "zero"},
			{V1: -2, V2: "two"},
			{V1: -4, V2: "four"},
		},
		helpers.CollectPairs(gi2.FilterMap(zeroToFour, evenNeg)),
	)

	assert.Equal(t,
		[]helpers.Pair[int, string]{
			{V1: 0, V2: "zero"},
			{V1: -2, V2: "two"},
		},
		helpers.CollectPairs(helpers.LimitPairs(2, gi2.FilterMap(zeroToFour, evenNeg))),
	)
}
