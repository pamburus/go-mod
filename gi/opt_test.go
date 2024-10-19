package gi_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi"
)

func TestIsNone(t *testing.T) {
	require.True(t, gi.IsNone(5, false))
	require.False(t, gi.IsNone(5, true))
}
