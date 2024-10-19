package gi2_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestEmpty(t *testing.T) {
	require.Equal(t, map[int]int{}, maps.Collect(gi2.Empty[int, int]()))
}
