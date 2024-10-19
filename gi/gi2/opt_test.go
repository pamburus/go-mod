package gi2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestIsNone(t *testing.T) {
	require.True(t, gi2.IsNone(5, "", false))
	require.False(t, gi2.IsNone(5, "", true))
}
