package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
)

func TestValues(t *testing.T) {
	require.Equal(t, []int(nil), slices.Collect(gi.Values[int]()))
	require.Equal(t, []int{42}, slices.Collect(gi.Values(42)))
	require.Equal(t, []int{42, 43}, slices.Collect(gi.Values(42, 43)))
}
