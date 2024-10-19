package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
)

func TestFind(t *testing.T) {
	values := slices.Values([]int{1, 2, 3, 4, 5})

	even := func(v int) bool {
		return v%2 == 0
	}

	found, ok := gi.Find(values, even)
	require.True(t, ok)
	require.Equal(t, 2, found)
}
