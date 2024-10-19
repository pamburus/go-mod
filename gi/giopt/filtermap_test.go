package giopt_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/giopt"
	"github.com/pamburus/go-mod/optional/optval"
)

func TestFilterMap(t *testing.T) {
	assert.Equal(t,
		[]int{60, 40},
		slices.Collect(giopt.FilterMap(slices.Values([]int{5, 3, 6, 7, 4}), func(v int) optval.Value[int] {
			return optval.New(v*10, v%2 == 0)
		})),
	)
}
