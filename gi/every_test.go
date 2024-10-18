package gi_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pamburus/go-mod/gi"
)

func TestEvery(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10, 11})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.False(t, result)
	})

	t.Run("AllTrue", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.True(t, result)
	})

	t.Run("AllFalse", func(t *testing.T) {
		values := slices.Values([]int{2, 4, 6, 8, 10, 11})

		predicate := func(v int) bool {
			return v%2 == 0
		}

		result := gi.Every(values, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		values := slices.Values([]int{})

		predicate := func(int) bool {
			return false
		}

		result := gi.Every(values, predicate)
		assert.True(t, result)
	})
}
