package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
	"github.com/pamburus/go-mod/gi/internal/testing/helpers"
)

func TestFilterMap(t *testing.T) {
	empty := slices.Values([]int{})
	oneToFour := slices.Values([]int{1, 2, 3, 4})

	evenX10 := func(v int) (int, bool) {
		if v%2 == 0 {
			return v * 10, true
		}

		return 0, false
	}

	require.Equal(t, []int(nil), slices.Collect(gi.FilterMap(empty, evenX10)))
	require.Equal(t, []int{20, 40}, slices.Collect(gi.FilterMap(oneToFour, evenX10)))
	require.Equal(t, []int{20}, slices.Collect(helpers.Limit(1, gi.FilterMap(oneToFour, evenX10))))
}
