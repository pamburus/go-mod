package giopt_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giopt"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestFind(t *testing.T) {
	assert.Equal(t,
		optval.New(6, true),
		giopt.Find(slices.Values([]int{5, 3, 6, 7, 4}), func(v int) bool {
			return v%2 == 0
		}),
	)
}
