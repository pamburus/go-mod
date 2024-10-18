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

func TestContainsLeft(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6})

		predicate := func(v1 int) bool {
			return v1 == 3
		}

		result := gi.ContainsLeft(pairs, predicate)
		assert.True(t, result)

		predicate = func(v1 int) bool {
			return v1 == 7
		}

		result = gi.ContainsLeft(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		pairs := slices.All([]int{})

		predicate := func(v1 int) bool {
			return v1 == 3
		}

		result := gi.ContainsLeft(pairs, predicate)
		assert.False(t, result)
	})
}

func TestContainsRight(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6})

		predicate := func(v2 int) bool {
			return v2 == 4
		}

		result := gi.ContainsRight(pairs, predicate)
		assert.True(t, result)

		predicate = func(v2 int) bool {
			return v2 == 7
		}

		result = gi.ContainsRight(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		pairs := slices.All([]int{})

		predicate := func(v2 int) bool {
			return v2 == 4
		}

		result := gi.ContainsRight(pairs, predicate)
		assert.False(t, result)
	})
}

func TestContainsKey(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6})

		predicate := func(v1 int) bool {
			return v1 == 3
		}

		result := gi.ContainsKey(pairs, predicate)
		assert.True(t, result)

		predicate = func(v1 int) bool {
			return v1 == 7
		}

		result = gi.ContainsKey(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		pairs := slices.All([]int{})

		predicate := func(v1 int) bool {
			return v1 == 3
		}

		result := gi.ContainsKey(pairs, predicate)
		assert.False(t, result)
	})
}

func TestContainsValue(t *testing.T) {
	t.Run("SomeTrue", func(t *testing.T) {
		pairs := slices.All([]int{1, 2, 3, 4, 5, 6})

		predicate := func(v2 int) bool {
			return v2 == 4
		}

		result := gi.ContainsValue(pairs, predicate)
		assert.True(t, result)

		predicate = func(v2 int) bool {
			return v2 == 7
		}

		result = gi.ContainsValue(pairs, predicate)
		assert.False(t, result)
	})

	t.Run("Empty", func(t *testing.T) {
		pairs := slices.All([]int{})

		predicate := func(v2 int) bool {
			return v2 == 4
		}

		result := gi.ContainsValue(pairs, predicate)
		assert.False(t, result)
	})
}
