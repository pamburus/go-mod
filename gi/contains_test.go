package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestContains(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		values := slices.Values([]int{1, 2, 3, 4, 5, 6})

		predicate := func(v int) bool {
			return v == 3
		}

		result := gi.Contains(values, predicate)
		assert.True(t, result)

		predicate = func(v int) bool {
			return v == 7
		}

		result = gi.Contains(values, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([]int{})

		predicate := func(v int) bool {
			return v == 3
		}

		result := gi.Contains(values, predicate)
		assert.False(t, result)
	})
}
