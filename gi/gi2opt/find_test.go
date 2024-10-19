package gi2opt_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi/gi2opt"
	"github.com/pamburus/go-mod/optional/optpair"
)

func TestFind(t *testing.T) {
	assert.Equal(t,
		optpair.New(2, 6, true),
		gi2opt.Find(slices.All([]int{5, 3, 6, 7, 4}), func(_, v int) bool {
			return v%2 == 0
		}),
	)
}
