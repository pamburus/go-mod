package gi2_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pamburus/go-mod/gi/gi2"
)

func TestSingle(t *testing.T) {
	require.Equal(t, map[int]string{42: "forty-two"}, maps.Collect(gi2.Single(42, "forty-two")))
}
