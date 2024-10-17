package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestFilterMap(t *testing.T) {
	empty := slices.Values([]int{})
	oneToFour := slices.Values([]int{1, 2, 3, 4})

	evenX10 := func(v int) optval.Value[int] {
		if v%2 == 0 {
			return optval.Some(v * 10)
		}

		return optval.None[int]()
	}

	require.Equal(t, []int(nil), slices.Collect(gi.FilterMap(empty, evenX10)))
	require.Equal(t, []int{20, 40}, slices.Collect(gi.FilterMap(oneToFour, evenX10)))
	require.Equal(t, []int{20}, slices.Collect(helpers.Limit(1, gi.FilterMap(oneToFour, evenX10))))
}
