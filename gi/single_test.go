package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestSingle(t *testing.T) {
	assert.Equal(t, []int{42}, slices.Collect(gi.Single(42)))
}
