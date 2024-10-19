package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
)

func TestEmpty(t *testing.T) {
	require.Equal(t, []int(nil), slices.Collect(gi.Empty[int]()))
}
